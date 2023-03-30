package api

import (
	"github.com/gorilla/mux"
	"github.com/timur-danilchenko/go-url-shortener/api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	// Route to create a new short URL
	router.HandleFunc("/shorten", handlers.CreateShortURL).Methods("POST")

	// Route to redirect to the original URL
	router.HandleFunc("/{id}", handlers.RedirectToURL).Methods("GET")
}
