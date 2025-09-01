package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

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

	GMAIL_USERNAME := os.Getenv("GMAIL_USERNAME")
	GMAIL_APP_PASSWORD := os.Getenv("GMAIL_APP_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", GMAIL_USERNAME) // change
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Ahooooy OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, GMAIL_USERNAME, GMAIL_APP_PASSWORD) // change

	return d.DialAndSend(m)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Serve static HTML form
	app.Static("/", "./public")

	// Handle registration form
	app.Post("/register", func(c *fiber.Ctx) error {
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
	})

	log.Fatal(app.Listen(":8080"))
}
