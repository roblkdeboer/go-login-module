package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/internal/db"
	"github.com/roblkdeboer/login-module/internal/handlers"
	"github.com/roblkdeboer/login-module/internal/middleware"

	_ "github.com/lib/pq"
)

func main() {

	db, err := db.ConnectDB("localhost", 5455,"postgresUser", "postgresPW", "postgresDB")
	if err != nil {
		log.Fatal(err)
	}

	// Create handlers with access to database connection
	userHandler := handlers.NewUserHandler(db)
	loginHandler := handlers.NewLoginHandler(db)

	router := mux.NewRouter()
	router.Use(middleware.RecoverMiddleware)

	// Apply auth middleware to a specific route
	router.Path("/admin").Methods(http.MethodGet).Handler(middleware.AuthMiddleware(http.HandlerFunc(handlers.AdminHandler)))


	userHandler.RegisterRoutes(router)
	loginHandler.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}