package main

import (
	"fmt"
	"net/http"
)

// RateLimitMiddleware wraps a handler and rejects requests that exceed the rate limit.
// Use the request's RemoteAddr as the clientID.
// Rejected requests should respond with HTTP 429 and a plain text message.
func RateLimitMiddleware(limiter *RateLimiter, next http.Handler) http.Handler {
	// TODO: implement
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.RemoteAddr
		if limiter.Allow(clientID) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Limit Exceeded", http.StatusTooManyRequests)
		}
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func setupServer(limiter *RateLimiter) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	return RateLimitMiddleware(limiter, mux)
}
