package warmup

// Write a test that exposes the off-by-one bug in Allow().
//
// Steps:
//   1. Create a RateLimiter with limit=3, window=1 minute
//   2. Call Allow("test-ip") exactly 3 times — all should return true
//   3. Call Allow("test-ip") a 4th time — this should return FALSE (rejected)
//      With the bug (> instead of >=), the 4th call returns TRUE — test fails
//      After fixing the bug (>= ), the 4th call returns FALSE — test passes
//
// This is the classic boundary test: verify that request N+1 is rejected
// when limit is N, not N+1.
