package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	//	"ahooooy/pkg/db"
	"ahooooy/pkg/mailer"
	"ahooooy/pkg/otp"
	"ahooooy/pkg/store"
	"ahooooy/service/registration/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "gorm.io/gorm"
)

var otpStore *redis.RedisOTPStore

//var dbConn *gorm.DB

func main() {

	rdb := store.InitRedis()
	otpStore = redis.NewRedisOTPStore(rdb)
	//	dbConn := db.SetupDB()
	ctx := context.Background()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.Static("/", "./public")

	app.Get("/", index)
	app.Post("/register", register(ctx))
	app.Post("/profile", saveProfile)

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

		enteredOtp := c.FormValue("otp")

		// Case 1: Validate OTP if provided
		if enteredOtp != "" {
			stored, err := otpStore.Get(ctx, email)
			if err != nil {
				log.Printf("❌ Failed to get OTP from Redis: %v", err)
				return c.Status(500).SendString("Failed to validate OTP")
			}
			if stored == nil {
				return c.Status(400).SendString("❌ OTP not found, please request again")
			}
			if stored.Code != enteredOtp {
				return c.Status(400).SendString("❌ Invalid OTP")
			}
			if time.Now().After(stored.ExpiresAt) {
				return c.Status(400).SendString("❌ OTP expired, please request again")
			}

			// ✅ OTP valid → registration successful
			log.Printf("🎉 Registered successfully: %s\n", email)
			return c.Redirect("/profile.html")
		}

		// Case 2: Generate and send OTP (existing flow)
		code := otp.Generate()
		log.Printf("📩 Generated OTP %s for email %s\n", code, email)

		otpData := redis.OTP{
			Email:     email,
			Code:      code,
			ExpiresAt: time.Now().Add(otp.OTPExpiry),
		}

		if err := otpStore.Set(ctx, otpData, otp.OTPExpiry); err != nil {
			log.Printf("❌ Failed to save OTP in Redis: %v", err)
			return c.Status(500).SendString("Failed to save OTP")
		}

		log.Println("✅ OTP stored in Redis")

		if err := mailer.SendEmail(email, code); err != nil {
			log.Printf("❌ Failed to send OTP email: %v", err)
			return c.Status(500).SendString("Failed to send OTP email")
		}

		return c.SendString("✅ OTP sent to " + email)
	}
}

// handler - save profile
func saveProfile(c *fiber.Ctx) error {
	name := c.FormValue("name")
	family := c.FormValue("family")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")

	if name == "" || family == "" || dob == "" || gender == "" {
		return c.Status(400).SendString("All fields are required")
	}

	// TODO: save to Redis / DB here (for now just log)
	log.Printf("👤 New profile created: %s %s, DOB: %s, Gender: %s",
		name, family, dob, gender)

	return c.SendString("🎉 Profile created successfully")
}
