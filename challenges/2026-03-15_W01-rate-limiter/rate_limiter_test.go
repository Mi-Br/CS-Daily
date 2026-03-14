package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// --- Unit tests: RateLimiter.Allow ---

func TestAllow_UnderLimit(t *testing.T) {
	rl := NewRateLimiter(3, time.Second)
	for i := 0; i < 3; i++ {
		if !rl.Allow("client1") {
			t.Fatalf("request %d should be allowed", i+1)
		}
	}
}

func TestAllow_AtLimit(t *testing.T) {
	rl := NewRateLimiter(3, time.Second)
	for i := 0; i < 3; i++ {
		rl.Allow("client1")
	}
	if rl.Allow("client1") {
		t.Fatal("4th request should be denied")
	}
}

func TestAllow_WindowReset(t *testing.T) {
	rl := NewRateLimiter(2, 50*time.Millisecond)

	rl.Allow("client1")
	rl.Allow("client1")

	if rl.Allow("client1") {
		t.Fatal("3rd request in window should be denied")
	}

	time.Sleep(60 * time.Millisecond) // wait for window to expire

	if !rl.Allow("client1") {
		t.Fatal("first request after window reset should be allowed")
	}
}

func TestAllow_IndependentClients(t *testing.T) {
	rl := NewRateLimiter(2, time.Second)

	rl.Allow("alice")
	rl.Allow("alice")

	// alice is at limit — bob should still be free
	if !rl.Allow("bob") {
		t.Fatal("bob should not be affected by alice's limit")
	}
	if rl.Allow("alice") {
		t.Fatal("alice should be denied (at limit)")
	}
}

func TestAllow_Concurrent(t *testing.T) {
	// Run with: go test -race
	const clients = 10
	const requests = 20

	rl := NewRateLimiter(5, time.Second)

	var wg sync.WaitGroup
	for c := 0; c < clients; c++ {
		clientID := fmt.Sprintf("client%d", c)
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			for i := 0; i < requests; i++ {
				rl.Allow(id)
			}
		}(clientID)
	}
	wg.Wait()
	// No panic = no data race (run with -race to verify)
}

// --- Integration tests: HTTP middleware ---

func TestMiddleware_AllowsUnderLimit(t *testing.T) {
	rl := NewRateLimiter(3, time.Second)
	server := setupServer(rl)

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)
		req.RemoteAddr = "1.2.3.4:1234"
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("request %d: expected 200, got %d", i+1, w.Code)
		}
	}
}

func TestMiddleware_Returns429WhenExceeded(t *testing.T) {
	rl := NewRateLimiter(2, time.Second)
	server := setupServer(rl)

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)
		req.RemoteAddr = "9.9.9.9:9999"
		httptest.NewRecorder()
		server.ServeHTTP(httptest.NewRecorder(), req)
	}

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.RemoteAddr = "9.9.9.9:9999"
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429, got %d", w.Code)
	}
}
