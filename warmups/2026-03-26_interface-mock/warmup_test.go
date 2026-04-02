package warmup

// Write 3 table-driven tests for RateLimitMiddleware using net/http/httptest.
//
// Test cases to cover:
//   1. MockLimiter{allowed: true}  → handler is called, response is 200
//   2. MockLimiter{allowed: false} → handler is NOT called, response is 429
//   3. MockLimiter{allowed: false} → body/headers confirm it was blocked (not just status)
//
// Hint: use httptest.NewRecorder() for the ResponseWriter and
//       httptest.NewRequest("GET", "/", nil) for the Request.
//
// Table structure suggestion:
//   name     string
//   allowed  bool
//   wantCode int
