package main

import (
	"fmt"
	"net/http"
	"os"

	"groupie-tracker/config"
	"groupie-tracker/handlers"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("too many arguments")
		return
	}
	err := config.InitTemplates("pages/*.html", "components/*.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artist", handlers.Artist)

	http.HandleFunc("/static/", handlers.Static)

	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
