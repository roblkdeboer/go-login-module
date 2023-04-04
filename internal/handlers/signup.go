package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/internal/errors"
	"github.com/roblkdeboer/login-module/internal/models"
	"github.com/roblkdeboer/login-module/internal/utils"
)

type UserHandler struct {
	DB *sql.DB
} 

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Use h.DB to interact with the database
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, errors.BadRequestError{Message: "Invalid request body"}.Error(), http.StatusBadRequest)
        return
    }

    // Check if the user already exists
    exists, err := utils.UserExists(h.DB, user.Email)
    if err != nil {
        dbErr := &errors.DatabaseError{Message: "Unable to query database"}
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
    }
    if exists {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("User already exists"))
        return
    }

    // Check if the password meets the required strength
    err = utils.IsValidPassword(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Generate password hash and salt
    passwordHash, err := utils.GeneratePasswordHash(user.Password)
    if err != nil {
        dbErr := &errors.DatabaseError{Message: "Error hashing the password"}
        http.Error(w, dbErr.Error(), http.StatusInternalServerError)
        return
    }

    err = utils.InsertUser(h.DB, user, passwordHash.Hash, passwordHash.Salt)
    if err != nil {
        dbErr := &errors.DatabaseError{Message: "Cannot insert into database"}
        http.Error(w, dbErr.Error(), http.StatusInternalServerError)
        return
    }
    
    // return a JSON response indicating success
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/sign-up", h.CreateUser).Methods("POST")
}