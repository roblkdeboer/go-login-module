package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/internal/errors"
	"github.com/roblkdeboer/login-module/internal/models"
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

    _, err = h.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		dbErr := &errors.DatabaseError{Message: err.Error()}
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

// func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
//     var user models.User
//     err := json.NewDecoder(r.Body).Decode(&user)
//     if err != nil {
//         http.Error(w, errors.BadRequestError{Message: "Invalid request body"}.Error(), http.StatusBadRequest)
//         return
//     }

//     // Insert new user record
// 	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
// 	if err != nil {
// 		dbErr := &errors.DatabaseError{Message: err.Error()}
// 		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

//     // return a JSON response indicating success
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
// }