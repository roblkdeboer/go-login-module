package utils

import (
	"database/sql"

	"github.com/roblkdeboer/login-module/internal/models"
)

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func InsertUser(db *sql.DB, user models.User, passwordHash string) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, passwordHash)
	return err
}

func UserExists(db *sql.DB, email string) (bool, error) {
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}