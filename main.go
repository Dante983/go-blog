package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dante983/go_one/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve database credentials from environment variables
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Ensure that the necessary environment variables are set
	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		log.Fatal("One or more database environment variables are not set")
	}

	// Construct the connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Connect to the MySQL database
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Pass the database connection to the route handlers
	routes.SetDB(db)

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define route handlers
	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/blog", routes.BlogHandler)
	http.HandleFunc("/about", routes.AboutHandler)
	http.HandleFunc("/create", routes.CreateHandler)
	http.HandleFunc("/update", routes.UpdateHandler)
	http.HandleFunc("/delete", routes.DeleteHandler)
	http.HandleFunc("/view", routes.ViewHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:80")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
