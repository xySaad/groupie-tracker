package main

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	var err error
	err = config.InitTemplates()
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist", handlers.Artist)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
