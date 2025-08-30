package redis

import "time"

// OTP stored in Redis with 30m TTL
// OTP represents a one-time passcode stored in Redis with a TTL.
type OTP struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`       // Always store as string (preserve leading zeros)
	ExpiresAt time.Time `json:"expires_at"` // For reference/debugging, Redis TTL enforces expiry
}
