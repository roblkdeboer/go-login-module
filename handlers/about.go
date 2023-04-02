package handlers

import (
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Learn more about us on the about page."))
}