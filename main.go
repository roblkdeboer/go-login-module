package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/roblkdeboer/login-module/handlers"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", handlers.HomeHandler)
    r.HandleFunc("/about", handlers.AboutHandler)

    log.Fatal(http.ListenAndServe("localhost:8080", r))
}