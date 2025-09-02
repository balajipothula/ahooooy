package mailer

import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

// sendEmail sends an email with the OTP
func SendEmail(to string, otp string) error {
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
