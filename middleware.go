package goralim

import (
	"fmt"
	"net/http"
)

func RateLimiter(tb *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.isAllowed() {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, "Too Many Requests")
			return
		}

		next.ServeHTTP(w, r)
	})
}

