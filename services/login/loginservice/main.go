package main

import (
    "dating-app/services/login/http"
    "dating-app/services/login/mysql"
    "dating-app/services/login/redis"
    "dating-app/services/login/service"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/redis/go-redis/v9"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // MySQL connection
    dsn := "user:password@tcp(localhost:3306)/datingdb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Redis connection
    rdb := redis2.NewClient(&redis2.Options{
        Addr: "localhost:6379",
    })

    // Setup Repos + Services
    userRepo := &mysql.UserRepository{DB: db}
    sessionRepo := &redis.SessionRepository{Client: rdb}
    authService := &service.AuthService{Users: userRepo, Sessions: sessionRepo}
    authHandler := &http.AuthHandler{Auth: authService}

    app := fiber.New()

    app.Post("/register", authHandler.Register)
    app.Post("/login", authHandler.Login)
    app.Post("/logout", authHandler.Logout)

    log.Println("Login service running on :3000")
    log.Fatal(app.Listen(":3000"))
}
