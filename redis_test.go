package goralim

import(
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewRedisClient(t *testing.T){
    // Test with valid config
    validConfig := RedisConfig{
        HOST: "localhost",
        PORT: 6379,
        PASS: "",
    }

    redisClientWithValidConfig := NewRedisClient(validConfig)
    assert.NotNil(t, redisClientWithValidConfig)

    _, err := redisClientWithValidConfig.Ping().Result()
    assert.NoError(t, err)

    // Test with invalid config
    invalidConfig := RedisConfig{
        HOST: "invalidhost",
        PORT: 1234,
        PASS: "",
    }

    redisClientWithWrongConfig := NewRedisClient(invalidConfig)
    assert.Nil(t, redisClientWithWrongConfig)

    // Test with nil config
    redisClientWithNilConfig := NewRedisClient(RedisConfig{})
    assert.Nil(t, redisClientWithNilConfig)

}
