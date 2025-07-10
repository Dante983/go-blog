// Run this script to generate a secure session key
// Usage: go run scripts/generate_session_key.go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// Generate 32 random bytes
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Failed to generate random key:", err)
	}

	// Encode to base64 for easy copying
	encoded := base64.StdEncoding.EncodeToString(key)

	fmt.Println("Generated Session Key:")
	fmt.Println(encoded)
	fmt.Println("\nAdd to your .env file:")
	fmt.Printf("SESSION_KEY=%s\n", encoded)
} 