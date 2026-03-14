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
// per `window` duration for each unique clientID.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	// TODO: implement
	return nil
}

// Allow returns true if the client is within their rate limit, false if exceeded.
// It must be safe to call concurrently from multiple goroutines.
func (r *RateLimiter) Allow(clientID string) bool {
	// TODO: implement
	return false
}
