// Run this script to generate a bcrypt hash for a password
// Usage: go run scripts/hash_password.go <password>
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run scripts/hash_password.go <password>")
		os.Exit(1)
	}

	password := os.Args[1]
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash: %s\n", hash)
	fmt.Printf("\nSQL to insert admin user:\n")
	fmt.Printf("INSERT INTO users (username, email, password_hash, is_admin) VALUES ('admin', 'admin@example.com', '%s', TRUE);\n", hash)
} 