package main

import (
	"log"
	"time"

	"ahooooy/service/registration/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Serve static form
	app.Static("/", "./public")

	// Handle registration
	app.Post("/register", func(c *fiber.Ctx) error {
		email := c.FormValue("email")

		// Simulate virtual number generator (e.g., internal/virtual/virtual_number.go)
		virtualNumber := "1234567890" // TODO: replace with generator

		// Create new member (not persisted yet, just demo)
		member := model.Member{
			VirtualNumber: virtualNumber,
			Email:         email,
			Verified:      false,
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
		}

		log.Printf("🆕 New member registered: %+v\n", member)

		return c.JSON(member) // return member as JSON for now
	})

	log.Fatal(app.Listen(":8080"))
}
