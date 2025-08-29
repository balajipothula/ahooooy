package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"dating-app/services/login/model"
)

type UserRepository struct {
	client *redis.Client
}

func NewUserRepository(client *redis.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) SaveUser(ctx context.Context, user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	// Store user JSON with email as key
	return r.client.Set(ctx, fmt.Sprintf("user:%s", user.Email), data, 0).Err()
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	val, err := r.client.Get(ctx, fmt.Sprintf("user:%s", email)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

