package main

import (
	"fmt"
	"testing"
)

// ### Exercise 5: Goroutines + Channels
// Write a function `FetchAll(urls []string) []Result` where:

// ```go
// type Result struct {
//     URL      string
//     Status   int
//     Err      error
// }
// ```

// Rules:
// - Fetch all URLs **concurrently** (use goroutines)
// - Use a real HTTP GET (doesn't matter if URLs are fake — handle the error gracefully)
// - Return results **in the same order as the input** (this is the tricky part)
// - Don't leak goroutines
// ---
func TestFetchAll(t *testing.T) {
	routes := []string{
		"sports",
		"workout",
		"ingredients",
		"users",
		"fakers",
		"makers",
		"not_found",
		"not_found2",
	}
	domain := "http://localhost"
	port := 6900
	urls := []string{}
	for _, r := range routes {
		urls = append(urls, fmt.Sprintf("%s:%d/%s", domain, port, r))
	}

	t.Run("Testing ability to fetch URLs", func(t *testing.T) {
		got := FetchAll(urls)
		if len(got) < len(urls) {
			t.Errorf("Not all results returned")
		}
	})
}
