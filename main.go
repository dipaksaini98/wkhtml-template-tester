package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	// Define the port to listen on
	port := ":8080"
	log.Printf("Server listening on port %s...\n", port)

	// Start the server
	log.Fatal(http.ListenAndServe(port, router))
}
