package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/internal/db"
	"github.com/roblkdeboer/login-module/internal/handlers"

	_ "github.com/lib/pq"
)

func main() {

	db, err := db.ConnectDB("localhost", 5455,"postgresUser", "postgresPW", "postgresDB")
	if err != nil {
		log.Fatal(err)
	}

	userHandler := handlers.NewUserHandler(db)

	router := mux.NewRouter()
	userHandler.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}