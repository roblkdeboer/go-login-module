package utils

import "database/sql"


func UserExists(db *sql.DB, email string) (bool, error) {
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}