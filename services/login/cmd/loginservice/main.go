package main

import (
	"log"

	localhttp "dating-app/services/login/http"
	"dating-app/services/login/mysql"
	redisrepo "dating-app/services/login/redis"
	"dating-app/services/login/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ----- DB -----
	dsn := "user:password@tcp(localhost:3306)/datingdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// ----- Redis -----
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// ----- Repos & Services -----
	userRepo := &mysql.UserRepository{DB: db}
	sessionRepo := &redisrepo.SessionRepository{Client: rdb}
	authService := &service.AuthService{Users: userRepo, Sessions: sessionRepo}

	// ----- Handlers -----
	local := &localhttp.AuthHandler{Auth: authService}
	oauth := &localhttp.OAuthHandler{Auth: authService}

	// ----- Fiber -----
	app := fiber.New()

	// Local auth
	app.Post("/register", local.Register)
	app.Post("/login", local.Login)
	app.Post("/logout", local.Logout)

	// OAuth routes
	app.Get("/auth/google/login", oauth.GoogleLogin)
	app.Get("/auth/google/callback", oauth.GoogleCallback)

	app.Get("/auth/facebook/login", oauth.FacebookLogin)
	app.Get("/auth/facebook/callback", oauth.FacebookCallback)

	app.Get("/auth/twitter/login", oauth.TwitterLogin)
	app.Get("/auth/twitter/callback", oauth.TwitterCallback)

	log.Println("Login service running on :3000")
	log.Fatal(app.Listen(":3000"))
}

