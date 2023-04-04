package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash represents the hash and salt of a password
type PasswordHash struct {
    Hash string
    Salt string
}

// GeneratePasswordHash generates a hash and salt for the given password
func GeneratePasswordHash(password string) (*PasswordHash, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    return &PasswordHash{Hash: string(hashedPassword)}, nil
}

// VerifyPassword checks if the given password matches the hash and salt
func VerifyPassword(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}