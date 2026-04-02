package warmup

// Write 3 tests for the functional options constructor.
//
// Test 1 — defaults:
//   - Call New() with no options
//   - Assert rl.limit == 100
//   - Assert rl.duration == 1*time.Minute
//   (You may need to export fields or add a helper to inspect them)
//
// Test 2 — custom limit + window:
//   - Call New(WithLimit(5), WithWindow(10*time.Second))
//   - Assert rl.limit == 5
//   - Assert rl.duration == 10*time.Second
//   - Verify Allow() enforces the custom limit (allow 5, reject 6th)
//
// Test 3 — cleanup goroutine fires (bonus):
//   - Call New(WithCleanupInterval(50*time.Millisecond))
//   - Inject some old timestamps directly into rl.requests
//   - time.Sleep(150*time.Millisecond)
//   - Assert the old timestamps have been evicted
//   Hint: Use a short interval so the test doesn't take long.
