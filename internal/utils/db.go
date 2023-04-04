package utils

import (
	"database/sql"

	"github.com/roblkdeboer/login-module/internal/models"
)


func UserExists(db *sql.DB, email string) (bool, error) {
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func InsertUser(db *sql.DB, user models.User, passwordHash, passwordSalt string) error {
	_, err := db.Exec("INSERT INTO users (name, email, password_hash, password_salt) VALUES ($1, $2, $3, $4)", user.Name, user.Email, passwordHash, passwordSalt)
	return err
}