package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"ahooooy/service/registration"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	DB_STR := os.Getenv("DB")
	DB_INT, err := strconv.Atoi(DB_STR)

	if err != nil {
		panic("Invalid DB index value: must be an integer")
	}

	// 1. Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ADDR"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DB:       DB_INT,
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
