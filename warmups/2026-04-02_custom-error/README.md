# 2026-04-02 — Custom Error Type

**Time:** ~20 min  
**Concept:** Custom error types, type assertions, HTTP headers

## Task

1. Implement a `RateLimitError` struct with:
   - `IP string`
   - `RetryAfter time.Duration`
   - `Error() string` method (implements the `error` interface)

2. Change `Allow()` signature to `Allow(ip string) (bool, error)`:
   - Returns `(false, &RateLimitError{...})` when rate limited
   - Returns `(true, nil)` when allowed

3. Update the middleware to:
   - Type-assert the error to `*RateLimitError`
   - Set the `Retry-After` header (value in seconds)
   - Return HTTP 429

## Why

Generic errors lose context. A typed error lets callers inspect fields (`RetryAfter`) programmatically — essential for the middleware to set the correct `Retry-After` header.

## Files

- `warmup.go` — `RateLimitError` struct, updated `Allow()` signature, middleware stub
- `warmup_test.go` — comment scaffold for 2 httptest tests
