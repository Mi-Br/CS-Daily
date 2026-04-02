# 2026-03-29 — Middleware Chaining

**Time:** ~20 min  
**Concept:** Higher-order functions, middleware composition, execution order

## Task

Implement `Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler`

The returned function should wrap the final handler with all middlewares applied **outermost-first**. That means the first middleware in the list runs first on the way in, and last on the way out.

## Example

```go
handler := Chain(logging, rateLimiting)(finalHandler)
// request flow: logging → rateLimiting → finalHandler
```

## Why

Real apps stack many middlewares: auth, logging, rate limiting, tracing. `Chain` lets you compose them cleanly without nesting.

## Files

- `warmup.go` — empty `Chain()` function with correct signature
- `warmup_test.go` — comment scaffold: chain 2 middlewares, verify execution order
