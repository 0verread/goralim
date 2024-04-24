package goralim

import(
	"time"
	"sync"

	redis "github.com/go-redis/redis"
)

// `Key` is the identifier for clint calls. This could be UserID/AccountID or Client IP
// `Capacity` is the max number of requests a client can make per second
// `RefillRate` is the rate bucket should be filled at to take new requests
// `Tokens` is the current number of tokens available in bucket
type TokenBucket struct {
	Key string
	RedisClient *redis.Client
	Capacity int
	RefillRate int
	Tokens int
	LastRefilledAt time.Time
	mutex sync.Mutex

}


func testFunc(){

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
	// Current time in milliseconds
	now := time.Now().UnixNano() / 1e6
	lastRefill, err := tb.RedisClient.Get(tb.Key + ":lastRefill").Int64()
	if err != nil {
		lastRefill = now
		tb.RedisClient.Set(tb.Key + ":lastRefill", now, 0)
	}
	elapsedTIme := now - lastRefill
	tokensToAdd := int(elapsedTIme) * tb.RefillRate / 1000
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




