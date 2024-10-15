package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/funcs"
)

func main() {
	// Handle root and artist pages
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/artist", funcs.ArtistHandler)

	// Serve static files like CSS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
