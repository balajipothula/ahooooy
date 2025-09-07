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
	ctx := context.Background()
	rdb := store.InitRedis()
	otpStore = redis.NewRedisOTPStore(rdb)

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(logger.New())

	// Serve static assets if needed
	app.Static("/assets", "./public/assets", fiber.Static{Browse: false})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	app.Post("/register", register(ctx))
	app.Post("/profile", saveProfile)

	// Run server
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Fiber stopped: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Printf("Fiber shutdown failed: %v", err)
	}
}

// Register handler
func register(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		if email == "" {
			return c.Status(400).SendString("Email is required")
		}

		enteredOtp := c.FormValue("otp")
		if enteredOtp != "" {
			stored, err := otpStore.Get(ctx, email)
			if err != nil {
				log.Printf("❌ Failed to get OTP: %v", err)
				return c.Status(500).SendString("Failed to validate OTP")
			}
			if stored == nil {
				return c.Status(400).SendString("❌ OTP not found")
			}
			if stored.Code != enteredOtp {
				return c.Status(400).SendString("❌ Invalid OTP")
			}
			if time.Now().After(stored.ExpiresAt) {
				return c.Status(400).SendString("❌ OTP expired")
			}

			log.Printf("🎉 OTP validated successfully: %s", email)
			return c.JSON(fiber.Map{
				"message": "OTP verified",
				"email":   email, // return email to client
			})
		}

		// Generate and send OTP
		code := otp.Generate()
		log.Printf("📩 Generated OTP %s for email %s", code, email)

		otpData := redis.OTP{
			Email:     email,
			Code:      code,
			ExpiresAt: time.Now().Add(otp.OTPExpiry),
		}
		if err := otpStore.Set(ctx, otpData, otp.OTPExpiry); err != nil {
			log.Printf("❌ Failed to save OTP: %v", err)
			return c.Status(500).SendString("Failed to save OTP")
		}

		if err := mailer.SendEmail(email, code); err != nil {
			log.Printf("❌ Failed to send OTP email: %v", err)
			return c.Status(500).SendString("Failed to send OTP email")
		}

		log.Println("✅ OTP stored and sent")
		return c.JSON(fiber.Map{
			"message": "OTP sent",
			"email":   email,
		})
	}
}

// Save profile handler
func saveProfile(c *fiber.Ctx) error {
	email := c.FormValue("email")
	firstName := c.FormValue("first_name")
	familyName := c.FormValue("family_name")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")

	if email == "" || firstName == "" || familyName == "" || dob == "" || gender == "" {
		return c.Status(400).SendString("All fields are required")
	}

	// TODO: insert into Postgres here
	log.Printf("👤 New profile created: %s %s, DOB: %s, Gender: %s, Email: %s",
		firstName, familyName, dob, gender, email)

	return c.JSON(fiber.Map{
		"message":     "Profile created successfully",
		"email":       email,
		"first_name":  firstName,
		"family_name": familyName,
		"dob":         dob,
		"gender":      gender,
	})
}
