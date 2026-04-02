package warmup

// Write 3 tests for AllowBurst().
//
// Test 1 — exactly n available (should pass):
//   - Create a RateLimiter with limit=5, fresh window
//   - Call AllowBurst("ip", 5) — should return true
//   - Call AllowBurst("ip", 1) again — should return false (no capacity left)
//
// Test 2 — n+1 needed (should fail):
//   - Create a RateLimiter with limit=3
//   - Call AllowBurst("ip", 4) — should return false (exceeds limit)
//   - Assert no tokens were consumed: Allow("ip") should still succeed 3 times
//
// Test 3 — edge case n=0:
//   - Create a RateLimiter with limit=0 (or any limit)
//   - Call AllowBurst("ip", 0) — should return true (zero always allowed)
//   - No tokens should be consumed
