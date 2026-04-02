# 2026-04-01 — Functional Options Pattern

**Time:** ~25 min  
**Concept:** Functional options, variadic constructors, clean API design

## Task

Refactor `New()` to use the functional options pattern:

```go
func New(opts ...Option) *RateLimiter
```

Implement these options:
- `WithLimit(n int) Option` — sets the request limit
- `WithWindow(d time.Duration) Option` — sets the time window
- `WithCleanupInterval(d time.Duration) Option` — (bonus) spawns a goroutine that periodically evicts old timestamps

## Why

Hard-coded constructor parameters become unwieldy as structs grow. Functional options let callers set only what they need, keep defaults for everything else, and add new options without breaking existing callers.

## Default values

```go
limit:    100
duration: 1 * time.Minute
```

## Files

- `warmup.go` — `Option` type, `New()` stub, option function stubs
- `warmup_test.go` — comment scaffold for 3 tests
