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

type LoginHandler struct {
	DB *sql.DB
} 

func (h *LoginHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	// Use h.DB to interact with the database
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errors.BadRequestError{Message: "Invalid request body"}.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve user information from the database
	dbUser, err := utils.GetUserByEmail(h.DB, user.Email)
	if err != nil {
		dbErr := &errors.DatabaseError{Message: "No email provided"}
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	if !utils.VerifyPassword(user.Password, dbUser.Password) {
		dbErr := &errors.DatabaseError{Message: "Invalid email or password"}
        http.Error(w, dbErr.Error(), http.StatusInternalServerError)
        return
	}

	// Create a JWT token and return it in the response
	tokenString, err := utils.GenerateJWTToken(dbUser.Email)
	if err != nil {
		http.Error(w, errors.BadRequestError{Message: "Authentication failed"}.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	json.NewEncoder(w).Encode(map[string]string{"message": "logged in successfully"})
}

func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{DB: db}
}

func (h *LoginHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.AuthenticateUser).Methods("POST")
}