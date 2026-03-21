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
	r.mu.Lock()
	defer r.mu.Unlock()
	client, exist := r.clients[clientID]
	if !exist {
		r.clients[clientID] = &clientState{
			count:       1,
			windowStart: time.Now(),
		}
		return true
	}
	if time.Since(client.windowStart) > r.window {
		client.count = 1
		client.windowStart = time.Now()
		return true

	} else if client.count < r.limit {
		client.count++
		return true
	}
	return false
}
