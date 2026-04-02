# 2026-03-26 — Interface + Mock for Testability

**Time:** ~20 min  
**Concept:** Interfaces, mock types, table-driven tests with httptest

## Task

1. Define a `Limiter` interface with a single method: `Allow(ip string) bool`
2. Update the HTTP middleware to accept `Limiter` instead of `*RateLimiter`
3. Implement a `MockLimiter` struct with an `allowed bool` field
4. Write 3 table-driven tests using `net/http/httptest` that cover: allowed request, blocked request, and correct status codes

## Why

Hard-coding `*RateLimiter` in the middleware makes it impossible to test without spinning up real timers. Accepting an interface lets you inject a `MockLimiter` in tests — no real rate limiting, just behavior you control.

## Files

- `warmup.go` — `Limiter` interface, `MockLimiter` struct, middleware stub
- `warmup_test.go` — comment scaffold for 3 table-driven tests
