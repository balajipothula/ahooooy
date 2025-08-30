package redis

import "time"

// OTP represents a one-time passcode stored in Redis with TTL.
// This is pure data — no Redis logic here.
type OTP struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`       // string to preserve leading zeros
	ExpiresAt time.Time `json:"expires_at"` // for debugging; Redis TTL enforces expiry
}
