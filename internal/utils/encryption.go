package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

// PasswordHash represents the hash and salt of a password
type PasswordHash struct {
    Hash string
    Salt string
}

// GeneratePasswordHash generates a hash and salt for the given password
func GeneratePasswordHash(password string) (*PasswordHash, error) {
    salt := make([]byte, 32)
    _, err := rand.Read(salt)
    if err != nil {
        return nil, err
    }

    hash := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)

    return &PasswordHash{Hash: hex.EncodeToString(hash), Salt: hex.EncodeToString(salt)}, nil
}

// VerifyPassword checks if the given password matches the hash and salt
func VerifyPassword(password string, hash *PasswordHash) bool {
    salt, _ := hex.DecodeString(hash.Salt)
    computedHash := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)
    return hex.EncodeToString(computedHash) == hash.Hash
}