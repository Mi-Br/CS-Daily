package warmup

import "net/http"

// Limiter is the interface the middleware depends on.
type Limiter interface {
	Allow(ip string) bool
}

// MockLimiter is a test double that always returns the same response.
type MockLimiter struct {
	allowed bool
}

// Allow implements the Limiter interface.
func (m *MockLimiter) Allow(ip string) bool {
	return false // TODO: return m.allowed
}

// RateLimitMiddleware wraps an http.Handler and enforces rate limiting.
// It accepts a Limiter so it can be tested with a MockLimiter.
func RateLimitMiddleware(l Limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: extract IP, call l.Allow(ip)
		// if not allowed: w.WriteHeader(http.StatusTooManyRequests) and return
		// otherwise: call next.ServeHTTP(w, r)
	})
}
