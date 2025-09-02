package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	gomail "gopkg.in/gomail.v2"
)

// generateOTP returns a random 6-digit OTP
func generateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n.Int64())
}

// sendEmail sends an email with the OTP
func sendEmail(to string, otp string) error {
	gmailUser := os.Getenv("GMAIL_USERNAME")
	gmailPass := os.Getenv("GMAIL_APP_PASSWORD")

	if gmailUser == "" || gmailPass == "" {
		return fmt.Errorf("missing Gmail credentials: check GMAIL_USERNAME and GMAIL_APP_PASSWORD env vars")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", gmailUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Ahooooy OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, gmailUser, gmailPass)

	return d.DialAndSend(m)
}

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
	otp := generateOTP()
	log.Printf("📩 Generated OTP %s for email %s\n", otp, email)

	// Send OTP via email
	if err := sendEmail(email, otp); err != nil {
		log.Printf("❌ Failed to send OTP email: %v", err)
		return c.Status(500).SendString("Failed to send OTP email")
	}

	return c.SendString("✅ OTP sent to " + email)
}
