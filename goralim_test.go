package goralim

import (
    "testing"

    "github.com/go-redis/redis"
)
func TestRefillTokens(t *testing.T){
    redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
    defer redisClient.Close()

    key := "test_key"
    capacity := 10
    refillRate := 5

    tb := NewTokenBucket(key, redisClient, capacity, refillRate)

    if tb.Key != key {
        t.Errorf("got %q, wanted %q", tb.Key, key)
    }
}


