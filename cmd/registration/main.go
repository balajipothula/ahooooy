package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ahooooy/pkg/mailer"
	"ahooooy/pkg/otp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.Get("/", index)
	app.Post("/register", register)

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
func register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "" {
		return c.Status(400).SendString("Email is required")
	}

	// Generate OTP
	otp := otp.Generate()

	log.Printf("📩 Generated OTP %s for email %s\n", otp, email)

	// Send OTP via email
	if err := mailer.SendEmail(email, otp); err != nil {
		log.Printf("❌ Failed to send OTP email: %v", err)
		return c.Status(500).SendString("Failed to send OTP email")
	}

	return c.SendString("✅ OTP sent to " + email)
}
