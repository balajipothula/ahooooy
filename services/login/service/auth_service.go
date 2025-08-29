package service

import (
	"context"
	"errors"
	"time"

	"dating-app/services/login/model"
	"dating-app/services/login/redisrepo"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users   *redisrepo.UserRepository
	jwtKey  []byte
	timeout time.Duration
}

func NewAuthService(users *redisrepo.UserRepository, jwtKey string) *AuthService {
	return &AuthService{
		users:   users,
		jwtKey:  []byte(jwtKey),
		timeout: time.Hour * 24, // token validity
	}
}

// Register a new user
func (s *AuthService) Register(ctx context.Context, email, password string) error {
	// check if user already exists
	existing, _ := s.users.GetUserByEmail(ctx, email)
	if existing != nil {
		return errors.New("user already exists")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// create user model
	user := &model.User{
		ID:           time.Now().UnixNano(), // crude unique ID
		Email:        email,
		PasswordHash: string(hash),
		Provider:     "local",
		CreatedAt:    time.Now(),
	}

	// save in Redis
	return s.users.SaveUser(ctx, user)
}

// Login validates credentials and returns JWT
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	// compare password hash
	if user.Provider == "local" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
			return "", errors.New("invalid credentials")
		}
	}

	// generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.timeout).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtKey)
}

// VerifyToken parses and validates JWT
func (s *AuthService) VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}

