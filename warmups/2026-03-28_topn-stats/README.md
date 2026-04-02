# 2026-03-28 — TopN() Sorted Stats Method

**Time:** ~20 min  
**Concept:** Sorting slices, struct methods, read locks

## Task

1. Implement `TopN(n int) []IPCount` on `RateLimiter`
2. It should return the top N IPs by in-window request count, sorted descending
3. Use a read lock — this is a read-only operation
4. If n is greater than the number of IPs, return all of them

## IPCount struct

```go
type IPCount struct {
    IP    string
    Count int
}
```

## Why

Observability. You want to know which IPs are hammering your service. `TopN` lets you surface the top offenders — useful for dashboards, alerts, or dynamic blocklists.

## Files

- `warmup.go` — `IPCount` struct, `RateLimiter` struct, empty `TopN()` stub
- `warmup_test.go` — comment scaffold for 2 tests
