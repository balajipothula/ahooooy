package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
