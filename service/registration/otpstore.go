package registration

import (
	rstore "ahooooy/service/registration/internal/redis"

	"github.com/redis/go-redis/v9"
)

// OTP alias (re-export) for convenience.
type OTP = rstore.OTP

// NewRedisOTPStore exposes the Redis-backed OTP store.
// This prevents external code from touching internal/ directly.
func NewRedisOTPStore(rdb *redis.Client) *rstore.RedisOTPStore {
	return rstore.NewRedisOTPStore(rdb)
}
