package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// OTP represents a one-time passcode stored in Redis.
// It is tied to an email and has a fixed expiry (enforced by Redis TTL).
type OTP struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`       // Always stored as string (preserve leading zeros like "000123")
	ExpiresAt time.Time `json:"expires_at"` // For reference/debugging, TTL in Redis is the actual source of truth
}

// RedisOTPStore manages OTP persistence in Redis.
// This is a thin helper layer around the Redis client to abstract away Redis operations.
type RedisOTPStore struct {
	client *redis.Client // Redis client connection
	ttl    time.Duration // Default TTL for OTPs (e.g., 30m)
}

// NewRedisOTPStore constructs a new RedisOTPStore with given client and TTL.
func NewRedisOTPStore(client *redis.Client, ttl time.Duration) *RedisOTPStore {
	return &RedisOTPStore{
		client: client,
		ttl:    ttl,
	}
}

// key builds the Redis key for storing OTPs.
// Example: otp:user@example.com
func (s *RedisOTPStore) key(email string) string {
	return "otp:" + email
}

// SetOTP stores an OTP in Redis with TTL.
// It marshals the OTP object into JSON and saves it as a value.
func (s *RedisOTPStore) SetOTP(ctx context.Context, otp OTP) error {
	data, err := json.Marshal(otp)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, s.key(otp.Email), data, s.ttl).Err()
}

// GetOTP retrieves the OTP for a given email from Redis.
// Returns nil if not found or expired.
func (s *RedisOTPStore) GetOTP(ctx context.Context, email string) (*OTP, error) {
	val, err := s.client.Get(ctx, s.key(email)).Result()
	if err == redis.Nil {
		return nil, nil // Not found or expired
	}
	if err != nil {
		return nil, err
	}

	var otp OTP
	if err := json.Unmarshal([]byte(val), &otp); err != nil {
		return nil, err
	}
	return &otp, nil
}

// DeleteOTP removes the OTP for the given email from Redis.
// Usually called after successful verification.
func (s *RedisOTPStore) DeleteOTP(ctx context.Context, email string) error {
	return s.client.Del(ctx, s.key(email)).Err()
}
