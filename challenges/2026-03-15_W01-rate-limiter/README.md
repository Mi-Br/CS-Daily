# Weekend Project W-01: Rate Limiter
**Track:** A + C  **ID:** W-01  **Date:** 2026-03-15

---

## What you're building

A **rate limiter** — one of the most common building blocks in real backend systems. Every API you've ever used has one. You're going to build it from scratch and wire it into an HTTP server as middleware.

**Goal:** No client should be able to make more than N requests per second. If they exceed it, respond with `HTTP 429 Too Many Requests`.

---

## The Concepts

This project connects the dots between what you've been learning:

| Concept | Where it shows up |
|---------|------------------|
| Maps | Track request counts per client |
| Structs | `RateLimiter`, `clientState` |
| Mutexes | Protect the map from concurrent writes |
| Time | Fixed-window expiry logic |
| HTTP middleware | Wrap a handler, intercept requests |

**Why mutex?** Your `Allow()` method will be called from multiple goroutines simultaneously (one per HTTP request). Without a mutex, two goroutines can read-then-write the same map entry at the same time — a data race. This is the "goroutines + shared state" problem.

---

## The Interfaces

Everything is stubbed out for you. You only need to implement:

### 1. `NewRateLimiter` — `rate_limiter.go`

```go
func NewRateLimiter(limit int, window time.Duration) *RateLimiter
```

Creates a limiter: max `limit` requests per `window` duration per client.

### 2. `Allow` — `rate_limiter.go`

```go
func (r *RateLimiter) Allow(clientID string) bool
```

Returns `true` if the client is within their limit, `false` if exceeded.

**Algorithm (fixed window):**
1. Lock the mutex
2. Look up the client in the map
3. If not found → create a new entry (count=1, windowStart=now) → return true
4. If found but window has expired → reset count to 1, reset windowStart → return true
5. If found and within window:
   - If count < limit → increment count → return true
   - If count >= limit → return false
6. Unlock the mutex

### 3. `RateLimitMiddleware` — `server.go`

```go
func RateLimitMiddleware(limiter *RateLimiter, next http.Handler) http.Handler
```

Wraps an HTTP handler. For each request:
- Extract client ID from `r.RemoteAddr`
- Call `limiter.Allow(clientID)`
- If allowed → call `next.ServeHTTP(w, r)`
- If denied → `http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)`

---

## Files

```
rate_limiter.go       ← implement NewRateLimiter + Allow
server.go             ← implement RateLimitMiddleware
rate_limiter_test.go  ← tests (already written — make them pass)
main.go               ← entry point (nothing to change here)
```

---

## Done When

- [ ] `go test ./...` passes
- [ ] `go test -race ./...` passes (no data races)
- [ ] `go vet ./...` clean
- [ ] You can run `go run .` and hit `localhost:8080/hello` — 6th request in a second gets 429

---

## Hints

<details>
<summary>Hint — mutex pattern</summary>

```go
func (r *RateLimiter) Allow(clientID string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    // ... safe to read/write r.clients here
}
```

`defer r.mu.Unlock()` ensures the lock is always released, even if you return early.
</details>

<details>
<summary>Hint — checking if a window has expired</summary>

```go
now := time.Now()
if now.Sub(state.windowStart) >= r.window {
    // window expired — reset
    state.count = 1
    state.windowStart = now
    return true
}
```
</details>

<details>
<summary>Hint — extracting IP from RemoteAddr</summary>

`r.RemoteAddr` is `"1.2.3.4:5678"`. You can use it as-is for the client ID (IP+port), or strip the port:

```go
host, _, _ := net.SplitHostPort(r.RemoteAddr)
```

Either works for this exercise.
</details>

---

## Stretch Goals (optional)

1. **Per-route limits** — make the limit configurable per route, not just global. Hint: you'll need a map of limiters or a way to tag requests.
2. **Sliding window** — instead of fixed window, implement a sliding window (harder, uses a queue of timestamps per client).
3. **Cleanup** — old client entries accumulate forever. Add a background goroutine that periodically deletes stale entries.

---

*Reviewed by Julia — your senior staff engineer on call.*
