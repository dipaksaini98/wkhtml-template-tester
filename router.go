package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router for your application.
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/", TemplateGenerator).Methods("GET")

	return router
}

// TemplateGenerator calls generate template function
func TemplateGenerator(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain")

	res, err := GenerateTemplate()

	if err != nil {
		// Write the response body on error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, res)
		fmt.Println(err)
		return
	}

	// Write the response body on success
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}
