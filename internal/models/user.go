package models

import (
	"database/sql"
	"time"

	"github.com/Dante983/go-blog/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	IsAdmin      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, is_admin, created_at, updated_at 
		FROM users 
		WHERE username = $1
	`

	var user User
	err := db.DB.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, is_admin, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`

	var user User
	err := db.DB.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user
func CreateUser(username, email, password string, isAdmin bool) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(`
		INSERT INTO users (username, email, password_hash, is_admin) 
		VALUES ($1, $2, $3, $4)`,
		username, email, hash, isAdmin,
	)

	return err
}
