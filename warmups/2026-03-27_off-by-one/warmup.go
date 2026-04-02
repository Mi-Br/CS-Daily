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

// Allow reports whether the given IP is within the rate limit.
// BUG: uses > instead of >= — allows one extra request per window.
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.duration)

	// Evict old timestamps
	times := rl.requests[ip]
	valid := times[:0]
	for _, t := range times {
		if t.After(windowStart) {
			valid = append(valid, t)
		}
	}
	rl.requests[ip] = valid

	count := len(valid)
	if count > rl.limit { // BUG: should be >=
		return false
	}

	rl.requests[ip] = append(rl.requests[ip], now)
	return true
}
