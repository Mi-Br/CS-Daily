# 2026-03-30 — Reset(ip) Write-Lock Method

**Time:** ~15 min  
**Concept:** Write locks, map mutation, no-op safety

## Task

Implement `Reset(ip string)` on `RateLimiter`:

1. Acquire a write lock (not a read lock — you're mutating)
2. Delete all recorded timestamps for the given IP
3. If the IP doesn't exist, do nothing (no-op, no error)

## Why

You might want to manually unblock a known-good IP — e.g., after a false positive, or when an admin clears a ban. `Reset` gives you that escape hatch. The write lock ensures no race with concurrent `Allow()` calls.

## Files

- `warmup.go` — `RateLimiter` struct, empty `Reset()` stub
- `warmup_test.go` — comment scaffold for 2 tests
