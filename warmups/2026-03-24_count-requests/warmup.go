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

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		duration: duration,
	}
}

// CountRequests returns the number of requests from ip that fall within the current window.
// It must not modify any state.
func (rl *RateLimiter) CountRequests(ip string) int {
	// your code here
	return 0
}
