package warmup

// Write 2 tests for the custom error middleware using net/http/httptest.
//
// Test 1 — 429 response has Retry-After header:
//   - Create a RateLimiter with limit=1
//   - Wrap a dummy handler with RateLimitMiddleware
//   - Send 2 requests to the same IP
//   - Assert the second response has status 429
//   - Assert the "Retry-After" header is set and is a positive integer string
//
// Test 2 — allowed request passes through:
//   - Create a RateLimiter with limit=5
//   - Wrap a handler that writes "ok" with status 200
//   - Send 1 request
//   - Assert response is 200
//   - Assert body is "ok"
//   - Assert no Retry-After header is present
