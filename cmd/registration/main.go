// cmd/registration/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ahooooy/service/registration/internal/redis"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// 1. Connect to Redis (adjust address if needed)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // default Redis
		Password: "",               // no password set
		DB:       9,                // use default DB
	})

	// 2. Wrap with our RedisOTPStore
	otpStore := redis.NewRedisOTPStore(rdb)

	// 3. Create a demo OTP
	otp := redis.OTP{
		Email:     "user@example.com",
		Code:      "123456", // string (preserves leading zeros)
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	if err := otpStore.Set(ctx, otp, 30*time.Minute); err != nil {
		log.Fatalf("failed to set OTP: %v", err)
	}
	fmt.Println("✅ OTP stored in Redis")

	// 4. Retrieve OTP
	stored, err := otpStore.Get(ctx, otp.Email)
	if err != nil {
		log.Fatalf("failed to get OTP: %v", err)
	}
	fmt.Printf("📩 Retrieved OTP for %s: %+v\n", otp.Email, stored)
}
