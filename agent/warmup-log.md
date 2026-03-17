# Daily Warmup Log

Track sent warmups so the cron agent doesn't repeat them.

## Format
`YYYY-MM-DD | topic | one-line description`

## Sent

<!-- entries added by cron agent after each send -->
2026-03-15 | mutex + map race condition | Spot the bug: map read outside lock in RateLimiter.Allow()
2026-03-16 | early returns refactor | Refactor nested canAccess() using guard clauses + write edge case tests
2026-03-17 | rate limiter cleanup | Implement cleanup() to remove expired timestamps and delete empty map keys
