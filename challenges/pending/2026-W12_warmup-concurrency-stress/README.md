# W12 Warmup — Concurrency Stress Test for RateLimiter

**Assigned:** 2026-03-25  
**Track:** W12 — Struct Design / Concurrency  
**Estimated time:** ~20 min  
**Package:** use the W-01 rate limiter code as base

---

## The Task

Prove your `RateLimiter` is actually safe under concurrent load by writing a stress test.

```go
func TestRateLimiter_ConcurrentAllow(t *testing.T) {
    rl := NewRateLimiter(10, time.Second) // 10 req/sec limit

    var wg sync.WaitGroup
    allowed := atomic.Int64{}
    rejected := atomic.Int64{}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // TODO: call rl.Allow("192.168.1.1")
            // and increment allowed or rejected accordingly
        }()
    }

    wg.Wait()

    t.Logf("allowed: %d, rejected: %d", allowed.Load(), rejected.Load())

    // TODO: write a meaningful assertion here
    // Hint: with limit=10 and 100 goroutines all firing at once,
    // what's the maximum number of allowed requests?
}
```

## Your Tasks

1. Fill in the goroutine body
2. Write a meaningful assertion — not just "it didn't crash"
3. Run with the race detector: `go test -race ./...`

## Questions to Think About

- With limit=10 and 100 goroutines firing simultaneously, what's the max allowed? Should your assertion be `== 10` or `<= 10`? Why?
- What if goroutines straddle a window boundary — could more than 10 slip through?
- What makes this test different from `TestAllow_Concurrent` you already wrote?

## Done When

- [ ] Test fills in the goroutine body correctly
- [ ] Assertion is meaningful (not just no-panic)
- [ ] Passes `go test -race ./...` clean
