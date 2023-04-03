package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/internal/handlers"
	"github.com/roblkdeboer/login-module/internal/middleware"
)

func main() {
    r := mux.NewRouter()

	// add the middleware
	r.Use(middleware.RecoverMiddleware)

    r.HandleFunc("/", handlers.HomeHandler)
    r.HandleFunc("/about", handlers.AboutHandler)
	r.HandleFunc("/sign-up", handlers.CreateUserHandler).Methods("POST")

    log.Fatal(http.ListenAndServe("localhost:8080", r))
}