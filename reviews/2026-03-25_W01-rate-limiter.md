# Review: W-01 Rate Limiter
**Date:** 2026-03-25  
**Challenge:** Maps, Mutexes, HTTP Middleware  
**Status:** ✅ Passed — all tests green, race detector clean

---

## What You Got Right

### Mutex placement — correct and complete
`sync.Mutex` on the struct, locked at the top of `Allow()`, deferred unlock. That's the right pattern and you got it right. The earlier commit (`fixed mutex lock`) shows you caught your own mistake — good instinct.

### Window reset logic — solid
```go
if time.Since(client.windowStart) > r.window {
    client.count = 1
    client.windowStart = time.Now()
    return true
}
```
Clean. Reset both `count` and `windowStart` atomically (within the lock). This is the correct fixed-window approach.

### Middleware — clean and minimal
```go
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    clientID := r.RemoteAddr
    if limiter.Allow(clientID) {
        next.ServeHTTP(w, r)
    } else {
        http.Error(w, "Limit Exceeded", http.StatusTooManyRequests)
    }
})
```
This is exactly right. You're using `http.HandlerFunc` as the adapter bridge (the pattern that tripped you up last time), calling `next.ServeHTTP(w, r)` to invoke the next handler, and using `http.Error(w, ..., http.StatusTooManyRequests)` correctly.

---

## Things to Note

### `RemoteAddr` includes the port
`r.RemoteAddr` returns `"1.2.3.4:5678"` — ip + port. For a rate limiter you usually want to key on IP only, not IP+port (each new connection from the same client gets a different ephemeral port). For this challenge it works because the tests use fixed `RemoteAddr` values, but in production you'd want:
```go
clientID, _, _ := net.SplitHostPort(r.RemoteAddr)
```
Worth knowing for the real world.

### `TODO` comment left in
```go
// TODO: implement
return http.HandlerFunc(...)
```
You implemented it — just left the comment behind. Tiny thing, but get in the habit of cleaning these up before pushing.

### Off-by-one edge case
When `count == limit` exactly, you return `false` — which means the limit is *exclusive* (you allow `limit - 1` requests, not `limit`). The test `TestAllow_AtLimit` sets limit=3, makes 3 calls, then checks the 4th is denied — so that test passes. But the spec says "up to `limit` requests" which implies `limit` should be allowed, not `limit - 1`.

Check: with `limit=3`, your code allows request 1 (`count=1`), 2 (`count=2`), 3... wait:
```go
} else if client.count < r.limit {  // 2 < 3 → true → count becomes 3
```
Actually on request 3: `count=2 < 3` → increments to 3, returns true. Then request 4: `count=3 < 3` → false. So you allow exactly 3 — that's correct! I misread. Ignore this point, logic is right.

---

## Summary

Clean, working, concurrent-safe rate limiter. The mutex fix you caught yourself was the main hurdle and you handled it. Middleware pattern is correct — you're getting comfortable with `http.HandlerFunc` and `ServeHTTP`.

**Score: ✅ Solid pass**

---

## What's Next

Moving to **A-02: Strings & Bytes** — more Go fundamentals, less net/http for now. I'll drop the challenge in `pending/` shortly.
