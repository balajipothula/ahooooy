package virtual

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// GenerateVirtualNumber creates a random 10-digit number starting with '9'.
// Example: 9XXXXXXXXX
func GenerateVirtualNumber() (string, error) {
	var sb strings.Builder
	sb.WriteString("9") // first digit fixed as '9'

	// generate remaining 9 digits securely
	for i := 0; i < 9; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("failed to generate digit: %w", err)
		}
		sb.WriteString(n.String())
	}

	return sb.String(), nil
}
