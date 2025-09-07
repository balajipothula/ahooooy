package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
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

	// Serve static assets only (CSS/JS/images)
	app.Static("/assets", "./public/assets", fiber.Static{Browse: false})

	app.Get("/", index)
	app.Post("/register", register(ctx))
	app.Get("/profile", profile)
	app.Post("/profile", saveProfile)

	// Optional: redirect accidental /profile.html requests
	app.Get("/profile.html", func(c *fiber.Ctx) error {
		email := c.Query("email")
		if email == "" {
			return c.Status(400).SendString("Missing email")
		}
		return c.Redirect("/profile?email=" + email)
	})

	// Run server
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

// Index handler
func index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
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

			log.Printf("🎉 Registered successfully: %s", email)
			return c.Redirect("/profile?email=" + email)
		}

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
		return c.SendString("✅ OTP sent to " + email)
	}
}

// Profile GET handler
func profile(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(400).SendString("Missing email")
	}

	htmlBytes, err := os.ReadFile("./public/profile.html")
	if err != nil {
		return c.Status(500).SendString("Could not load profile.html")
	}

	html := strings.Replace(string(htmlBytes), "{{.Email}}", email, 1)
	return c.Type("html").SendString(html)
}

// Save profile POST handler
func saveProfile(c *fiber.Ctx) error {
	email := c.FormValue("email")
	firstName := c.FormValue("first_name")
	familyName := c.FormValue("family_name")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")

	if email == "" || firstName == "" || familyName == "" || dob == "" || gender == "" {
		return c.Status(400).SendString("All fields are required")
	}

	log.Printf("👤 New profile created: %s %s, DOB: %s, Gender: %s, Email: %s",
		firstName, familyName, dob, gender, email)

	return c.SendString("🎉 Profile created successfully")
}
