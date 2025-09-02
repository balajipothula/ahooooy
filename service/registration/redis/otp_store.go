package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisOTPStore handles OTP persistence in Redis.
type RedisOTPStore struct {
	rdb *redis.Client
}

// NewRedisOTPStore creates a new instance backed by Redis.
func NewRedisOTPStore(rdb *redis.Client) *RedisOTPStore {
	return &RedisOTPStore{rdb: rdb}
}

func (s *RedisOTPStore) Set(ctx context.Context, otp OTP, ttl time.Duration) error {
	key := "otp:" + otp.Email
	data, err := json.Marshal(otp)
	if err != nil {
		return err
	}
	return s.rdb.Set(ctx, key, data, ttl).Err()
}

func (s *RedisOTPStore) Get(ctx context.Context, email string) (*OTP, error) {
	key := "otp:" + email
	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var otp OTP
	if err := json.Unmarshal([]byte(val), &otp); err != nil {
		return nil, err
	}
	return &otp, nil
}

func (s *RedisOTPStore) Delete(ctx context.Context, email string) error {
	key := "otp:" + email
	return s.rdb.Del(ctx, key).Err()
}

// Verify checks if the provided code matches the stored OTP.
func (s *RedisOTPStore) Verify(ctx context.Context, email, code string) (bool, error) {
	otp, err := s.Get(ctx, email)
	if err != nil {
		return false, err
	}
	return otp.Code == code, nil
}
