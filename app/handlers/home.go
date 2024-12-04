package handlers

import (
	"groupie-tracker/config"
	"groupie-tracker/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var artists []utils.Object

	err := utils.FetchData(BaseURL+"/artists", &artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	err = config.Templates.ExecuteTemplate(w, "home.html", artists)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
