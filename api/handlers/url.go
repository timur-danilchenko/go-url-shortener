package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"github.com/timur-danilchenko/go-url-shortener/models"
)

var urls []models.URL

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a short ID for the URL
	id, err := shortid.Generate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the URL to the list of URLs
	url.ID = id
	urls = append(urls, url)

	// Send the short URL back to the client
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		fmt.Sprintf("http://%s/%s", r.Host, id),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectToURL(w http.ResponseWriter, r *http.Request) {
	// Find the URL with the specified ID
	vars := mux.Vars(r)
	id := vars["id"]
	var url models.URL
	for _, u := range urls {
		if u.ID == id {
			url = u
			break
		}
	}

	// Redirect to the original URL
	http.Redirect(w, r, url.URL, http.StatusSeeOther)
}
