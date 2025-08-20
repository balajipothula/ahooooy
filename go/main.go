package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Song struct {
	SongID     uint      `gorm:"column:songId;primaryKey;autoIncrement" json:"songId,omitempty"`
	Artist     string    `gorm:"column:artist;not null" json:"artist"`
	Title      string    `gorm:"column:title;not null" json:"title"`
	Difficulty float32   `gorm:"column:difficulty;not null" json:"difficulty"`
	Level      int16     `gorm:"column:level;not null;check:level > 0 AND level < 100" json:"level"`
	Released   time.Time `gorm:"column:released;type:date;not null" json:"released"`
}

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	/*
		app.Use(logger.New())
		app.Use(logger.New(logger.Config{
			Output: io.Discard,
		}))
	*/
	app.Get("/", index)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			//log.Printf("app server stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//fmt.Println("\nshutting down app server...")

	// Gracefully shutdown Fiber
	if err := app.Shutdown(); err != nil {
		//log.Fatalf("app shutdown failed: %v", err)
	}

}

// Handler - Index
func index(c *fiber.Ctx) error {
	return c.SendString("🪶 Feathery Fast APIs with 🐹 GO Fiber")
}
