package main

import (
	"fmt"
	"net/http"

	goralim "github.com/0verread/goralim"
)

func main() {
	config := goralim.RedisConfig {
		HOST: "127.0.0.1",
		PORT: 6379,
		AUTH: "",
	}
	redisStore := goralim.NewRedisClient(config)
	tb := goralim.NewTokenBucket("usr-ty789", redisStore, 10, 5)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hey There this is a Request")
	})

	rateLimitedHandler := goralim.RateLimiter(tb, handler)
	http.ListenAndServe(":8080", rateLimitedHandler)
}

