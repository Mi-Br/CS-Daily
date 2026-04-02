package warmup

import "net/http"

// Chain composes multiple middlewares into a single middleware.
// Middlewares are applied outermost-first: the first in the list
// wraps the outermost layer and runs first on every request.
func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	// TODO: iterate middlewares in reverse order, wrapping handler each time
	// Hint: start from the innermost (last) and work outward
	return nil
}
