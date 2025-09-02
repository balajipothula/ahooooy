package store

import (
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisDBStr := os.Getenv("REDIS_DB")
	redisDBInt, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Fatalf("❌ Invalid REDIS_DB value %q: must be an integer", redisDBStr)
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDBInt,
	})
}
