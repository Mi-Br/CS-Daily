package warmup

// Write a test that verifies execution order of chained middlewares.
//
// Setup:
//   - Create a []string log (shared via closure or pointer)
//   - Define middlewareA: appends "A" to the log before calling next
//   - Define middlewareB: appends "B" to the log before calling next
//   - Define a final handler: appends "handler" to the log
//
// Test:
//   - Call Chain(middlewareA, middlewareB)(finalHandler)
//   - Serve a test request using httptest.NewRecorder() + httptest.NewRequest()
//   - Assert the log is ["A", "B", "handler"] — outermost-first
//
// Bonus: test with 3 middlewares and verify full order is preserved.
