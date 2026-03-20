package main

import (
	"sync"
	"time"
)

type clientState struct {
	count       int
	windowStart time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	clients map[string]*clientState
}

// NewRateLimiter creates a RateLimiter that allows up to `limit` requests
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]*clientState),
	}

}

// Allow returns true if the client is within their rate limit, false if exceeded.
// It must be safe to call concurrently from multiple goroutines.
func (r *RateLimiter) Allow(clientID string) bool {
	client, exist := r.clients[clientID]
	if !exist {
		r.mu.Lock()
		r.clients[clientID] = &clientState{
			count:       1,
			windowStart: time.Now(),
		}
		r.mu.Unlock()
		return true
	}
	if time.Since(client.windowStart) > r.window {
		r.mu.Lock()
		client.count = 1
		client.windowStart = time.Now()
		r.mu.Unlock()
		return true

	} else if client.count < r.limit {
		r.mu.Lock()
		client.count++
		r.mu.Unlock()
		return true
	}
	return false
}
