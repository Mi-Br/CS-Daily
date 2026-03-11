package main

import (
	"fmt"
	"net/http"
	"sync"
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

type Result struct {
	URL    string
	Status int
	Err    error
}

func FetchAll(urls []string) []Result {
	res := make([]Result, len(urls))
	var wg sync.WaitGroup
	for i, u := range urls {
		fmt.Println(i)
		wg.Add(1)

		go func(i int, u string) {
			defer wg.Done()
			r, err := http.Get(u)
			if err != nil {
				defer r.Body.Close()
				res[i] = Result{Status: 400, Err: err}
			}
			res[i] = Result{Status: r.StatusCode, Err: err}
		}(i, u)
	}
	wg.Wait()
	return res
}
