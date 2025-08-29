package redis

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type SessionRepository struct {
    Client *redis.Client
}

func (r *SessionRepository) StoreToken(token string, userID uint, expiration time.Duration) error {
    ctx := context.Background()
    return r.Client.Set(ctx, token, userID, expiration).Err()
}

func (r *SessionRepository) ValidateToken(token string) (string, error) {
    ctx := context.Background()
    return r.Client.Get(ctx, token).Result()
}

func (r *SessionRepository) RevokeToken(token string) error {
    ctx := context.Background()
    return r.Client.Del(ctx, token).Err()
}
