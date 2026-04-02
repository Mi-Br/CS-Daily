package warmup

import (
	"net/http"
	"sync"
	"time"
)

// RateLimitError is returned by Allow() when a request is rate limited.
type RateLimitError struct {
	IP         string
	RetryAfter time.Duration
}

// Error implements the error interface.
func (e *RateLimitError) Error() string {
	// TODO: return a human-readable message, e.g.:
	// fmt.Sprintf("rate limit exceeded for %s, retry after %s", e.IP, e.RetryAfter)
	return ""
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

// Allow reports whether the request from ip is within the rate limit.
// Returns (true, nil) if allowed, or (false, *RateLimitError) if rate limited.
func (rl *RateLimiter) Allow(ip string) (bool, error) {
	// TODO: implement sliding window logic
	// TODO: on rejection, return false, &RateLimitError{IP: ip, RetryAfter: ...}
	return false, nil
}

// RateLimitMiddleware wraps next, returning 429 with Retry-After header when limited.
func RateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: extract IP from r.RemoteAddr
		// TODO: call rl.Allow(ip)
		// TODO: if error, type-assert to *RateLimitError
		//       set Retry-After header (seconds), WriteHeader(429), return
		// TODO: otherwise call next.ServeHTTP(w, r)
	})
}
