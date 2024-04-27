package goralim

import (
	"fmt"
	"net/http"

  "github.com/gin-gonic/gin"
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

func ginRateLimiter(tb *TokenBucket) gin.HandlerFunc {
  return func (c *gin.ContextA)  {
    if !tb.isAllowed(){
      c.AbortWithStatus(http.StatusTooManyRequests)
      return
    }
    c.Next()
  }
}

