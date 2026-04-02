package warmup

// Write 2 tests for Reset().
//
// Test 1 — mid-window reset:
//   - Create a RateLimiter with limit=2
//   - Call Allow("1.2.3.4") twice — both succeed, now at limit
//   - Call Allow("1.2.3.4") — should return false (at limit)
//   - Call Reset("1.2.3.4")
//   - Call Allow("1.2.3.4") again — should return true (counter cleared)
//
// Test 2 — unknown IP no-op:
//   - Create a RateLimiter
//   - Call Reset("unknown-ip") — should not panic, no error
//   - Assert Allow("unknown-ip") still works normally after the no-op reset
