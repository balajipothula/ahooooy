package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"ahooooy/service/registration"

	"github.com/redis/go-redis/v9"
)

// initRedis initializes and returns a Redis client using env vars.
func initRedis() *redis.Client {
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

func main() {
	ctx := context.Background()

	// 1. Setup Redis
	rdb := initRedis()
	otpStore := registration.NewRedisOTPStore(rdb)

	// 2. Example usage (demo flow)
	otp := registration.OTP{
		Email:     "user@example.com",
		Code:      "123456",
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	if err := otpStore.Set(ctx, otp, 30*time.Minute); err != nil {
		log.Fatalf("❌ failed to set OTP: %v", err)
	}
	log.Println("✅ OTP stored in Redis")

	stored, err := otpStore.Get(ctx, otp.Email)
	if err != nil {
		log.Fatalf("❌ failed to get OTP: %v", err)
	}
	log.Printf("📩 Retrieved OTP for %s: %+v\n", otp.Email, stored)
}
