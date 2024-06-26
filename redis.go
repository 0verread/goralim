package goralim

import (
	"fmt"
	"log"

	redis "github.com/go-redis/redis"
)

type RedisConfig struct {
	HOST string
	PORT int
	PASS string
}

func NewRedisClient(config RedisConfig) *redis.Client {

    // TODO: remove fmt print
    addr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	client := redis.NewClient(&redis.Options {
		Addr: addr,
		Password: config.PASS,
		PoolSize: 200,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Println("Failed to connect with Redis", err)
		client = nil
	}

	return client
}


