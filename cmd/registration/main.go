package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ahooooy/pkg/mailer"
	"ahooooy/pkg/otp"
	"ahooooy/pkg/store"
	"ahooooy/service/registration/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var otpStore *redis.RedisOTPStore

func main() {

	rdb := store.InitRedis()
	otpStore = redis.NewRedisOTPStore(rdb)
	ctx := context.Background()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.Get("/", index)
	app.Post("/register", register(ctx))

	// Run server in goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("fiber stopped: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Printf("fiber shutdown failed: %v", err)
	}

}

// handler - index
func index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
}

// handler - register
func register(ctx context.Context) fiber.Handler {

	return func(c *fiber.Ctx) error {

		email := c.FormValue("email")

		if email == "" {
			return c.Status(400).SendString("Email is required")
		}

		// Generate OTP
		code := otp.Generate()
		log.Printf("📩 Generated OTP %s for email %s\n", code, email)

		// Create OTP struct (consistent with earlier demo flow)
		otpData := redis.OTP{
			Email:     email,
			Code:      code,
			ExpiresAt: time.Now().Add(otp.OTPExpiry),
		}

		// Save OTP in Redis
		if err := otpStore.Set(ctx, otpData, otp.OTPExpiry); err != nil {
			log.Printf("❌ Failed to save OTP in Redis: %v", err)
			return c.Status(500).SendString("Failed to save OTP")
		}

		log.Println("✅ OTP stored in Redis")

		stored, err := otpStore.Get(ctx, email)
		if err != nil {
			log.Fatalf("❌ failed to get OTP: %v", err)
		}

		log.Printf("📩 Retrieved OTP for %s: %+v\n", email, *stored)

		// Send OTP via email
		if err := mailer.SendEmail(email, code); err != nil {
			log.Printf("❌ Failed to send OTP email: %v", err)
			return c.Status(500).SendString("Failed to send OTP email")
		}

		return c.SendString("✅ OTP sent to " + email)

	}

}
