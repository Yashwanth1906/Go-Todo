package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect() {
	log.Println("Starting database connection...")

	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		log.Println("Continuing without .env file...")
	} else {
		log.Println(".env file loaded successfully")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is empty")
	}

	log.Printf("Connecting to database with DSN: %s", dsn)
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		panic(err)
	}

	log.Println("Testing database connection...")
	if err = database.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		log.Fatal(err)
	}

	log.Println("Database connection successful!")
	db = database
}

func GetDB() *sql.DB {
	return db
}

// func SetDB(dsn *gorm.DB) {
// 	db = dsn
// }
