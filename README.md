![banner](images/banner.png)

# Goralim

[goralim](https://github.com/0verread/goralim) is a Golang package that provides a rate limiter based on [Token bucket](https://en.wikipedia.org/wiki/Token_bucket) algorithm. It is capabale to handle distributed workload with its redis database support. It has HTTP server middleware support (as of now).

> 🚧 this is a beta version now and under active development. For production use, fork it and made changes based on your need.

## Install
```bash
go get "github.com/0verread/goralim" -m
```
## Usage

```golang
package main

import (
	"fmt"
	"net/http"

	goralim "github.com/0verread/goralim"
)

func main() {
	// Redis initialization
	config := goralim.RedisConfig{
		HOST: "127.0.0.1",
		PORT: 6379,
		// password
		PASS: "",
	}
	redisStore := goralim.NewRedisClient(config)

	// setup rate limiter for a key (key can be userId/client)
	// 10 is bucket size, 5 is bucket refill rate per second
	tb := goralim.NewTokenBucket("key123", redisStore, 10, 5)

	// create HTTP server and setup a goralim middleware to put a rate limiter
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hey There this is a Request")
	})

	rateLimitedHandler := goralim.RateLimiter(tb, handler)
	http.ListenAndServe(":8080", rateLimitedHandler)
}
```

## License
Under [MIT](LICENSE) license
