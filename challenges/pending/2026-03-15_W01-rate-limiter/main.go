package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	limiter := NewRateLimiter(5, time.Second)
	server := setupServer(limiter)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
