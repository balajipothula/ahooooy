package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generate returns a random 6-digit OTP
func Generate() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n.Int64())
}
