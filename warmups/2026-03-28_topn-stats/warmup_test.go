package warmup

// Write 2 tests for TopN().
//
// Test 1 — basic sort order:
//   - Create a RateLimiter
//   - Simulate different IPs with different request counts (inject timestamps directly
//     into rl.requests, or call Allow() repeatedly with different IPs)
//   - Call TopN(2)
//   - Assert the result is sorted descending by Count
//   - Assert the top IP is the one with the most requests
//
// Test 2 — n > actual number of IPs:
//   - Create a RateLimiter with only 2 tracked IPs
//   - Call TopN(10)
//   - Assert the result has length 2 (not 10, not a panic)
//   - Assert both IPs are returned
