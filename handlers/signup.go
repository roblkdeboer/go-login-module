package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roblkdeboer/login-module/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // TODO: save the user to a database or other storage mechanism

    // return a JSON response indicating success
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}