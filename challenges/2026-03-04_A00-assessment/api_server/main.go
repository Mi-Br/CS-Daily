package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

func routeHandler(route string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		delay := time.Second * 1 * time.Duration(rand.IntN(10))
		time.Sleep(delay)
		fmt.Fprintf(w, "Route: %s", route)
	}
}

func main() {

	urls := []string{
		"sports",
		"workout",
		"ingredients",
		"users",
		"fakers",
		"makers",
	}

	fmt.Print("Starting api server")

	for _, u := range urls {

		path := fmt.Sprintf("/%s", u)
		http.HandleFunc(path, routeHandler(u))
	}

	http.ListenAndServe(":6900", nil)
}
