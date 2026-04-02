# 2026-03-31 — AllowBurst(ip, n) Atomic Burst

**Time:** ~20 min  
**Concept:** Atomic operations under a single lock, all-or-nothing semantics

## Task

Implement `AllowBurst(ip string, n int) bool` on `RateLimiter`:

- Returns `true` only if **n tokens are available** in the current window
- If true, **consumes all n** under a single lock (atomic)
- If false, consumes nothing
- Edge case: `n=0` always returns `true`

## Why

Some clients need to reserve capacity in bulk — e.g., a batch job sending N notifications at once. All-or-nothing semantics prevent partial consumption that could confuse callers. The single-lock guarantee means no TOCTOU race.

## Files

- `warmup.go` — `RateLimiter` struct, empty `AllowBurst()` stub
- `warmup_test.go` — comment scaffold for 3 tests
