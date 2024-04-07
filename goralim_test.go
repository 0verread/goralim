package goralim

import (
    "testing"

    "github.com/go-redis/redis"
    "github.com/stretchr/testify/assert"
)

func setup() *TokenBucket{
    redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
    defer redisClient.Close()

    key := "test_key"
    capacity := 10
    refillRate := 5

    tb := NewTokenBucket(key, redisClient, capacity, refillRate)

    return tb
}

func TestRefillTokens(t *testing.T){
    tb := setup()
    expected_value := "test_key"
    assert.Equal(t, tb.Key, expected_value)
}

func TestIsAllowed(t *testing.T){
    tb := setup()
    for i:=0; i < 10; i++ {
        tb.isAllowed()
    }

    assert.False(t, tb.isAllowed())

}


