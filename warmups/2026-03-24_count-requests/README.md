# Warmup: CountRequests
**Date:** 2026-03-24 | **Est. time:** ~20 min  
**Prereading:** `reading/2026-03-23_mutexes-rwmutex.md` ← read this first

---

## Context

This warmup uses a **sliding window** `RateLimiter` — different from the fixed-window
one in W-01. Instead of storing a count + timestamp, it stores a slice of request timestamps
per IP. That lets you count exactly how many requests fall within the current window.

```go
type RateLimiter struct {
    mu       sync.RWMutex
    requests map[string][]time.Time
    limit    int
    duration time.Duration
}
```

---

## The Task

Implement a **read-only** method that returns how many valid (within-window) requests
an IP has made — without modifying any state.

```go
func (rl *RateLimiter) CountRequests(ip string) int {
    // your code here
}
```

**Rules:**
- Use the right kind of lock — you're not writing anything
- Count only timestamps still within the current window (`time.Since(ts) <= rl.duration`)
- Return 0 if the IP has no entries

---

## Write two tests

Put them in `warmup_test.go`. No scaffolding — write them from scratch.

1. **IP with mixed timestamps** — some expired, some valid. Assert only the valid ones are counted.
2. **Unknown IP** — an IP not in the map at all. Assert the result is 0.

---

## Done when
- [ ] `CountRequests` implemented in `warmup.go`
- [ ] Both tests written in `warmup_test.go`
- [ ] `go test ./...` passes
- [ ] `go test -race ./...` passes
