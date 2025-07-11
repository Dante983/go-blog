package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// First, check if we have a DATABASE_URL (common for cloud providers)
	databaseURL := os.Getenv("DATABASE_URL")

	var dsn string

	if databaseURL != "" {
		// Use the DATABASE_URL directly
		dsn = databaseURL
		log.Println("Using DATABASE_URL for connection")
	} else {
		// Fall back to individual parameters
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
		sslMode := os.Getenv("DB_SSL_MODE")

		// Default values
		if port == "" {
			port = "5432"
		}
		if sslMode == "" {
			sslMode = "disable"
		}

		// PostgreSQL connection string format
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, pass, name, sslMode)
		log.Println("Using individual parameters for connection")
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected!")
}
