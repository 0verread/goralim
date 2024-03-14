package goralim

import(
	"time"
	"sync"

	redis "github.com/go-redis/redis"
)

type TokenBucket struct {
	Key string
	RedisClient *redis.Client
	Capacity int
	RefillRate int
	Tokens int
	LastRefilledAt time.Time
	mutex sync.Mutex

}

func NewTokenBucket(key string, redisClient *redis.Client, capacity int, refillRate int) *TokenBucket {
	return &TokenBucket {
		Key: key,
		RedisClient: redisClient,
		Capacity: capacity,
		RefillRate: refillRate,
		Tokens: capacity,
		LastRefilledAt: time.Now(),
	}
}

func (tb *TokenBucket) refillTokens() {
	now := time.Now().UnixNano() / 1e6
	lastRefill, err := tb.RedisClient.Get(tb.Key + ":lastRefill").Int64()
	if err != nil {
		lastRefill = now
		tb.RedisClient.Set(tb.Key + ":lastRefill", now, 0)
	}
	elapsedTIme := now - lastRefill
	tokensToAdd := int(elapsedTIme) * (tb.RefillRate/1000)

	currentTokens, _ := tb.RedisClient.Get(tb.Key).Int()
	newTokens := currentTokens + tokensToAdd
	if newTokens > tb.Capacity {
		newTokens = tb.Capacity
	}
	tb.RedisClient.Set(tb.Key, newTokens, 0)
	tb.RedisClient.Set(tb.Key + ":lastRefill", now, 0)
}

func (tb *TokenBucket) isAllowed() bool {
	tb.refillTokens()
	currentTokens, _ := tb.RedisClient.Get(tb.Key).Int()
	
	if currentTokens > 0 {
		tb.RedisClient.Decr(tb.Key)
		return true
	}
	return false
}
