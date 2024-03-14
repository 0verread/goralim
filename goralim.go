package goralim

import(
	"time"
	"sync"
	"math"
)

type TokenBucket struct {
	Key string
	Capacity int
	RefillRate float64 //Bucket refil rate
	Tokens int
	lastRefilledAt time.Time
	mutex sync.Mutex

}

func NewTokenBucket(key string, capacity int, refillRate float64) *TokenBucket {
	return &TokenBucket {
		Key: key,
		Capacity: capacity,
		RefillRate: refillRate,
		Tokens: capacity,
		lastRefilledAt: time.Now(),
	}
}

func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	elapsedTIme := now.Sub(tb.lastRefilledAt).Seconds()
	tokensToAdd := float64(elapsedTIme)*tb.RefillRate
	tb.Tokens = int(math.Min(tokensToAdd+float64(tb.Tokens), float64(tb.Capacity)))
	tb.lastRefilledAt = now
}

func (tb *TokenBucket) isAllowed() bool {
	tb.refillTokens()
	if tb.Tokens > 0 {
		tb.Tokens -= 1
		return true
	}
	return false
}
