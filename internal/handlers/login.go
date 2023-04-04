package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginHandler struct {
	DB *sql.DB
} 

func (h *LoginHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Authentication"))
}

func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{DB: db}
}

func (h *LoginHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.AuthenticateUser).Methods("POST")
}