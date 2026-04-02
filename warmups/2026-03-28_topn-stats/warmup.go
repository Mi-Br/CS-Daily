package warmup

import (
	"sync"
	"time"
)

// IPCount holds an IP address and its in-window request count.
type IPCount struct {
	IP    string
	Count int
}

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

// TopN returns the top n IPs by in-window request count, sorted descending.
// If n > number of tracked IPs, all IPs are returned.
func (rl *RateLimiter) TopN(n int) []IPCount {
	// TODO: acquire read lock
	// TODO: for each IP, count timestamps within the current window
	// TODO: sort by count descending
	// TODO: return first n (or all if fewer than n)
	return nil
}
