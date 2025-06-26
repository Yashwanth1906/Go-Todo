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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	db = database
}

func GetDB() *sql.DB {
	return db
}

// func SetDB(dsn *gorm.DB) {
// 	db = dsn
// }
