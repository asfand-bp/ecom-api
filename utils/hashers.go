package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given plain text password using bcrypt.
func HashPassword(password string) (string, error) {
	// Generate a salted hash for the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided plain text password matches the hashed password.
func VerifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
