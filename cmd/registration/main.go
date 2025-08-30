package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ahooooy/service/registration"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// 1. Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       9,
	})

	// 2. Use exposed wrapper
	otpStore := registration.NewRedisOTPStore(rdb)

	// 3. Create a demo OTP
	otp := registration.OTP{
		Email:     "user@example.com",
		Code:      "123456",
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
