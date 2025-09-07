package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings" // ✅ needed for placeholder replacement
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

	// Serve only static assets (CSS/JS/images), NOT profile.html
	app.Static("/assets", "./public/assets", fiber.Static{
		Browse: false,
	})

	app.Get("/", index)
	app.Post("/register", register(ctx))
	app.Get("/profile", profile) // ✅ inject email
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
			return c.Redirect("/profile?email=" + email)
		}

		// Case 2: Generate and send OTP
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

// handler - profile (GET) → inject email into profile.html
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

// handler - save profile (POST)
func saveProfile(c *fiber.Ctx) error {
	email := c.FormValue("email")
	first_name := c.FormValue("first_name")
	family_name := c.FormValue("family_name")
	dob := c.FormValue("dob")
	gender := c.FormValue("gender")

	if email == "" || first_name == "" || family_name == "" || dob == "" || gender == "" {
		return c.Status(400).SendString("All fields are required")
	}

	// TODO: save to Redis / DB here (for now just log)
	log.Printf("👤 New profile created: %s %s, DOB: %s, Gender: %s, Email: %s",
		first_name, family_name, dob, gender, email)

	return c.SendString("🎉 Profile created successfully")
}
