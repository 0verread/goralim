package main

import (
	"fmt"
	"net/http"

	"goralim"
)

func main() {
	tb := goralim.NewTokenBucket("usr-1234", 10, 5)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hey There this is a Request")
	})

	rateLimitedHandler := goralim.RateLimiter(tb, handler)
	http.ListenAndServe(":8080", rateLimitedHandler)
}