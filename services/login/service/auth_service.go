package service

import (
    "dating-app/services/login/model"
    "dating-app/services/login/mysql"
    "dating-app/services/login/redis"
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

type AuthService struct {
    Users    *mysql.UserRepository
    Sessions *redis.SessionRepository
}

func (s *AuthService) Register(email, password string) (*model.User, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Email:        email,
        PasswordHash: string(hash),
        Provider:     "local",
    }

    if err := s.Users.Create(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
    user, err := s.Users.FindByEmail(email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    claims := jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(15 * time.Minute).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    if err := s.Sessions.StoreToken(signed, user.ID, 15*time.Minute); err != nil {
        return "", err
    }

    return signed, nil
}

func (s *AuthService) Logout(token string) error {
    return s.Sessions.RevokeToken(token)
}
