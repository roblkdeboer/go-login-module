package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// Claims represents the data we want to store in the JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWTToken generates a JWT for the given email
func GenerateJWTToken(email string) (string, error) {
	// Create the claims
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("error signing token")
	}

	return tokenString, nil
}