package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

type Song struct {
	Id         uint    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Artist     string  `gorm:"column:artist;not null" json:"artist"`
	Title      string  `gorm:"column:title;not null" json:"title"`
	Difficulty float32 `gorm:"column:difficulty;not null" json:"difficulty"`
	Level      int16   `gorm:"column:level;not null;check:level > 0 AND level < 100" json:"level"`
	Released   string  `gorm:"column:released;type:date;not null" json:"released"`
}

func (Song) TableName() string {
	return `music_schema."Song"`
}

func GetPostgresDSN() string {

	host := os.Getenv("SUPABASE_PG_HOST")
	port := os.Getenv("SUPABASE_PG_PORT")
	user := os.Getenv("SUPABASE_PG_USER")
	password := os.Getenv("SUPABASE_PG_PASSWORD")
	dbname := os.Getenv("SUPABASE_PG_DB_NAME")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("Missing required Supabase Postgres environment variables")
	}

	// Construct DSN string
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)
}

func setupDB() *gorm.DB {

	dsn := GetPostgresDSN()

	// Parse pgx config
	pgxConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("failed to parse pgx config: %v", err)
	}

	// 🚀 Force pgx to always use simple protocol (no prepared statements)
	pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Open stdlib DB with pgx config
	sqlDB := stdlib.OpenDB(*pgxConfig)

	// Connect with GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		PrepareStmt: false,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}

func main() {

	database = setupDB()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", index)
	app.Post("/insert/song", insertSong)
	app.Get("/select/songs", selectSongs)
	app.Get("/select/song/:id", selectSongById)
	app.Put("/update/song/:id", updateSongById)
	app.Delete("/delete/song/:id", deleteSongById)

	// Run server in goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("fiber stopped: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Printf("fiber shutdown failed: %v", err)
	}
}

// handler - index
func index(c *fiber.Ctx) error {
	return c.SendString("🪶 Feathery Fast APIs with 🐹 GO Fiber")
}

// handler - insert song
func insertSong(c *fiber.Ctx) error {
	jsonSong := new(Song)

	if err := c.BodyParser(jsonSong); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json song record"})
	}

	if jsonSong.Level <= 0 || jsonSong.Level >= 100 {
		return c.Status(400).JSON(fiber.Map{"error": "level must be between 1 and 99"})
	}

	if jsonSong.Artist == "" || jsonSong.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "song artist and title are required"})
	}

	if err := database.Create(jsonSong).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "unable to insert song record"})
	}

	return c.Status(201).JSON(jsonSong)
}

// handler - select songs
func selectSongs(c *fiber.Ctx) error {

	var songs []Song

	if err := database.Find(&songs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch songs"})
	}

	return c.JSON(songs)

}

// handler - select song by id
func selectSongById(c *fiber.Ctx) error {

	id := c.Params("id")
	var song Song

	if err := database.First(&song, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "song not found"})
	}

	return c.JSON(song)

}

// handler - update song by id
func updateSongById(c *fiber.Ctx) error {

	id := c.Params("id")
	var song Song

	if err := database.First(&song, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "song not found"})
	}

	jsonSong := new(Song)
	if err := c.BodyParser(jsonSong); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid json song record"})
	}

	// JSON Song `level` value validation.
	if jsonSong.Level <= 0 || jsonSong.Level >= 100 {
		return c.Status(400).JSON(fiber.Map{"error": "song level must be between 1 and 99"})
	}
	// JSON Song `artist` and `title` values validation.
	if jsonSong.Artist == "" || jsonSong.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "song artist and title are required"})
	}

	// Copying JSON Song values into song variable.
	song.Artist = jsonSong.Artist
	song.Title = jsonSong.Title
	song.Difficulty = jsonSong.Difficulty
	song.Level = jsonSong.Level
	song.Released = jsonSong.Released

	if err := database.Save(&song).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update song"})
	}

	return c.JSON(song)

}

// handler - delete song by id
func deleteSongById(c *fiber.Ctx) error {

	id := c.Params("id")
	var song Song

	if err := database.First(&song, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "song not found"})
	}

	if err := database.Delete(&song).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete song"})
	}

	return c.SendStatus(204)

}
