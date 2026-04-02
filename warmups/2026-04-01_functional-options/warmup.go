package warmup

import (
	"sync"
	"time"
)

// Option is a functional option for configuring a RateLimiter.
type Option func(*RateLimiter)

type RateLimiter struct {
	mu       sync.RWMutex
	requests map[string][]time.Time
	limit    int
	duration time.Duration
}

// New creates a RateLimiter with sensible defaults, modified by any provided Options.
// Default: limit=100, window=1 minute.
func New(opts ...Option) *RateLimiter {
	// TODO: create rl with default values (limit=100, duration=1*time.Minute)
	// TODO: apply each option: for _, opt := range opts { opt(rl) }
	// TODO: return rl
	return nil
}

// WithLimit returns an Option that sets the maximum requests per window.
func WithLimit(n int) Option {
	return nil // TODO: return func(rl *RateLimiter) { rl.limit = n }
}

// WithWindow returns an Option that sets the time window duration.
func WithWindow(d time.Duration) Option {
	return nil // TODO: return func(rl *RateLimiter) { rl.duration = d }
}

// WithCleanupInterval returns an Option that spawns a background goroutine
// which evicts expired timestamps at the given interval.
func WithCleanupInterval(d time.Duration) Option {
	return nil // TODO: spawn a goroutine that ticks at interval d and evicts old entries
}
