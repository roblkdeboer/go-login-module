package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Protected Page"})
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin", AdminHandler).Methods("POST")
}