package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roblkdeboer/login-module/internal/db"
	"github.com/roblkdeboer/login-module/internal/errors"
	"github.com/roblkdeboer/login-module/internal/models"

	_ "github.com/lib/pq"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, errors.BadRequestError{Message: "Invalid request body"}.Error(), http.StatusBadRequest)
        return
    }

    // Connect to database
	db, err := db.Connect()
	if err != nil {
		dbErr := &errors.DatabaseError{Message: err.Error()}
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

    // TODO: save the user to a database or other storage mechanism

    // return a JSON response indicating success
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}