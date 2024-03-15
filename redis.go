package goralim

import (
	"fmt"
	"log"

	redis "github.com/go-redis/redis"
)

type RedisConfig struct {
	HOST string
	PORT int
	AUTH string
}

func NewRedisClient(config RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	client := redis.NewClient(&redis.Options {
		Addr: addr,
		Password: config.AUTH,
		PoolSize: 200,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Println("Failed to connect with Redis", err)
		client = nil
	}

	return client
}
