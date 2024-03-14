package main

import (
	"fmt"
	"net/http"

	goralim "github.com/0verread/goralim"
)

func main() {
	redisStore := goralim.NewRedisClient(&goralim.RedisConfig {
		HOST: "127.0.0.1",
		PORT: 6379,
		AUTH: "",
	})
	tb := goralim.NewTokenBucket("usr-1234", redisStore, 10, 5)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hey There this is a Request")
	})

	rateLimitedHandler := goralim.RateLimiter(tb, handler)
	http.ListenAndServe(":8080", rateLimitedHandler)
}

