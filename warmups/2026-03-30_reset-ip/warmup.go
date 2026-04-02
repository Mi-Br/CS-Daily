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

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.duration)

	times := rl.requests[ip]
	valid := times[:0]
	for _, t := range times {
		if t.After(windowStart) {
			valid = append(valid, t)
		}
	}
	rl.requests[ip] = valid

	if len(valid) >= rl.limit {
		return false
	}

	rl.requests[ip] = append(rl.requests[ip], now)
	return true
}

// Reset clears all recorded timestamps for the given IP.
// If the IP is not tracked, this is a no-op.
func (rl *RateLimiter) Reset(ip string) {
	// TODO: acquire write lock
	// TODO: delete rl.requests[ip]
}
