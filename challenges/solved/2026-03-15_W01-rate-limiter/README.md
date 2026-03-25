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

---

## 🔒 Advanced Extension — Edge API Gateway Rate Limiter

> **Prerequisites before attempting this:**
> - W-01 base challenge complete ✅
> - Comfortable with Redis (or willing to learn)
> - Understand HTTP middleware patterns
> - Know what a distributed system is and why shared state is hard

This is a real-world design & implementation challenge based on production systems. Take it on when you're confident in the basics.

---

### The Problem

Extend your rate limiter to work as middleware in an **Edge API Gateway** — the entry point that handles all traffic before it reaches your services.

The gateway needs to enforce policies like:
- Allow only **10 req/hr** per `request.body.username` on `/login`
- Allow only **5 req/15min** per `request.body.password` on `/login`
- Allow only **10 req/min** per `request.headers.hostip` globally

This is fundamentally different from W-01: you're no longer rate-limiting by IP. You're rate-limiting by **arbitrary request properties** with **per-policy rules**.

---

### Design Questions to Answer First

Before writing a single line of code, think through these. Write your answers in a `DESIGN.md` file.

**Storage: In-Memory vs Redis**

Your W-01 solution uses in-memory state. That works for one server. What happens with 10 servers?

- In-memory + sticky routing: consistent hashing sends the same key to the same node. What breaks when a node restarts? When you scale up?
- Redis: shared state across all nodes. `INCR` is atomic. But now Redis is on the critical path — what's your plan if Redis goes down? Fail open (allow all) or fail closed (block all)?
- Hybrid: local in-memory cache + Redis as source of truth. Accept ~5% over-counting in exchange for fewer Redis calls. Is that acceptable for your use case?

**Policy Extensibility**

Hardcoding policies is a dead end. Design a declarative policy schema:

```json
{
  "path": "/login",
  "extract": "body.username",
  "limit": 10,
  "window": "1h",
  "action": "block"
}
```

Think about:
- How do you safely extract nested keys from request body/headers? (dot-notation parser, handle missing fields)
- If the extracted value is null/missing — do you skip enforcement or apply a global fallback bucket?
- Can you rate-limit on a composite key? (`username + ip` together)
- What happens when multiple policies match the same request? Priority rules?

**Correctness vs Scale**

| Granularity | Accuracy | Memory cost | Use case |
|-------------|----------|-------------|----------|
| Per-second  | High     | High        | DDoS protection |
| Per-minute  | Good     | Medium      | API abuse |
| Per-hour    | Approximate | Low      | Business quotas |

At high throughput, exact counting is expensive. Discuss the tradeoff:
- Is it acceptable to allow 10% overage if it means 10x throughput?
- Probabilistic structures like Count-Min Sketch give O(1) space with bounded error — worth it?

**Fixed Window vs Sliding Window**

You built a fixed window in W-01. The boundary burst problem: 10 requests at 13:59 + 10 at 14:01 = 20 requests in 2 minutes while technically within limits.

Three approaches:
1. **Fixed window** — simple, O(1), has boundary burst flaw
2. **Sliding window log** — store timestamps per client in a sorted set, count `[now - window, now]`. Accurate, but O(N) memory per client
3. **Sliding window counter** (best tradeoff) — keep current + previous window, weighted estimate:
   ```
   count = prev_count × (1 - elapsed/window) + curr_count
   ```
   O(1) memory, ~1% error, Redis-friendly

---

### Implementation Tasks

Once you've answered the design questions:

**Phase 1 — Policy Engine**
- [ ] Define a `Policy` struct with `path`, `extract`, `limit`, `window`, `action`
- [ ] Write a dot-notation extractor that safely pulls values from request body/headers
- [ ] Load policies from a JSON/YAML config file at startup
- [ ] Match incoming requests to policies

**Phase 2 — Redis Backend**
- [ ] Replace the in-memory map with Redis using `INCR` + `EXPIRE` for fixed window
- [ ] Implement the sliding window counter algorithm in Redis (two keys per window)
- [ ] Add a circuit breaker: if Redis is unreachable, fail open with a warning log

**Phase 3 — Security**
- [ ] `request.body.password` as a Redis key is a critical security flaw — hash all sensitive extracted values (SHA-256) before using them as keys
- [ ] Add `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` response headers
- [ ] Return `Retry-After` header on 429

**Phase 4 — Production Hardening**
- [ ] Hot-reload policies without restarting (file watcher or config endpoint)
- [ ] Background cleanup goroutine to remove stale Redis keys
- [ ] Benchmark: what's your throughput at 1k, 10k, 100k clients?

---

### Done When

- [ ] Policies are driven by config, not hardcoded
- [ ] Works correctly across multiple server instances (Redis-backed)
- [ ] Sensitive keys are hashed before storage
- [ ] `DESIGN.md` explains your tradeoff decisions
- [ ] Load test shows no race conditions under concurrent traffic

---

### Further Reading (when you're ready)

- [Cloudflare: How we built rate limiting](https://blog.cloudflare.com/counting-things-a-lot-of-different-things/)
- [Redis pattern: Sliding window rate limiter](https://redis.io/glossary/rate-limiting/)
- [Token bucket vs leaky bucket vs sliding window](https://stripe.com/blog/rate-limiters)
