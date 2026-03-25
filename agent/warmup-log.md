# Daily Warmup Log

Track sent warmups so the cron agent doesn't repeat them.

## Format
`YYYY-MM-DD | topic | one-line description`

## Sent

<!-- entries added by cron agent after each send -->
2026-03-15 | mutex + map race condition | Spot the bug: map read outside lock in RateLimiter.Allow()
2026-03-16 | early returns refactor | Refactor nested canAccess() using guard clauses + write edge case tests
2026-03-17 | rate limiter cleanup | Implement cleanup() to remove expired timestamps and delete empty map keys
2026-03-18 | HTTP middleware wrapper | Implement withRateLimit() middleware — extract IP, call Allow(), return 429 or pass through
2026-03-19 | table-driven tests | Write 5 table-driven test cases for Allow() covering edge cases + window expiry
2026-03-20 | sync.RWMutex | Explain RLock vs Lock: why Count uses RLock but Allow needs Lock, and what breaks if you mix them up
2026-03-21 | missing return in middleware | Spot the bug: no early return after WriteHeader in withRateLimit(), causing next() to still run on rejected requests
2026-03-22 | CountRequests read-only method | Implement CountRequests(ip) using RLock, counting only in-window timestamps
2026-03-23 | Stats() snapshot method | Implement Stats() returning IP→count map using RLock, no mutation, empty map if no in-window requests
2026-03-24 | struct-based refactor | Refactor package-level vars into a RateLimiter struct with constructor and method receivers
