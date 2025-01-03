package handlers

import (
	"bytes"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, "404 - page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		Error(w, "405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var artists []utils.Object

	err := utils.FetchData(BaseURL+"/artists", &artists)
	if err != nil {
		Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	err = config.Templates.ExecuteTemplate(&buffer, "home.html", artists)
	if err != nil {
		Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	buffer.WriteTo(w)
}
