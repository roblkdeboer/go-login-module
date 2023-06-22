package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/roblkdeboer/login-module/internal/db"
	"github.com/roblkdeboer/login-module/internal/handlers"
	"github.com/roblkdeboer/login-module/internal/middleware"
)

func main() {

	err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatal("error loading .env file", err)
    }

	portNum := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portNum)
	if err != nil {
		log.Fatal("Error converting PORT to integer:", err)
	}
	
	db, err := db.ConnectDB("localhost", port,os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
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