# 2026-03-27 — Off-by-One in Allow()

**Time:** ~15 min  
**Concept:** Bug hunting, boundary conditions, test-driven debugging

## Task

1. Read the `Allow()` method in `warmup.go` — there's a subtle bug on the limit check
2. Write a test that **fails** against the buggy code (i.e., it exposes the bug)
3. Fix the bug so the test passes

## The Bug

```go
if count > rl.limit {  // wrong: allows one extra request
```

Should be:

```go
if count >= rl.limit {  // correct: rejects at exactly the limit
```

## Why

Off-by-one errors in rate limiters mean real users get one free extra request per window. In abuse scenarios, that multiplies. Tests that verify the exact boundary are essential.

## Files

- `warmup.go` — intentionally buggy `Allow()` using `>`
- `warmup_test.go` — comment scaffold: write the test that catches it
