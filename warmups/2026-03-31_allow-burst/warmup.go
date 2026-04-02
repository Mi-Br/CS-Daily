package warmup

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	limit    int
	duration time.Duration
}

func New(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		duration: duration,
	}
}

// AllowBurst reports whether n tokens are available for the given IP,
// and if so, consumes all n atomically under a single lock.
// Returns false without consuming anything if fewer than n slots remain.
// n=0 always returns true.
func (rl *RateLimiter) AllowBurst(ip string, n int) bool {
	// TODO: acquire write lock (single lock for check + consume)
	// TODO: evict timestamps outside the current window
	// TODO: check if len(valid) + n <= rl.limit
	// TODO: if yes, append n timestamps and return true
	// TODO: if no, return false (consume nothing)
	return false
}
