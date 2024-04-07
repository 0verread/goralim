package goralim

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/go-redis/redis"
    "github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T){
    redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
    defer redisClient.Close()

    key := "test_key"
    capacity := 10
    refillRate := 5

    tb := NewTokenBucket(key, redisClient, capacity, refillRate)

    handler :=  http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
            w.WriteHeader(http.StatusOK)
    })

    rateLimiter := RateLimiter(tb, handler)

    req, err := http.NewRequest("GET", "/", nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    rateLimiter.ServeHTTP(rr, req)


    // Test tokens exhaustion
    for i := 0; i < capacity; i++ {
        req, err := http.NewRequest("GET", "/", nil)
        assert.NoError(t, err)

        rr := httptest.NewRecorder()
        rateLimiter.ServeHTTP(rr, req)

    }

    req, err = http.NewRequest("GET", "/", nil)
    assert.NoError(t, err)

    rr = httptest.NewRecorder()
    rateLimiter.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusTooManyRequests, rr.Code)
    assert.Contains(t, rr.Body.String(), "Too Many Requests")
}


