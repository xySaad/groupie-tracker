package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		http.Error(w, "404 - page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "501 - method not implemented", http.StatusMethodNotAllowed)
		return
	}

	queries := r.URL.Query()

	artistID, msg, status := utils.ValidateArtistID(queries)

	if status != http.StatusOK {
		http.Error(w, msg, status)
		return
	}

	artist, err := utils.FetchArtist(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	relation, err := utils.FetchArtistRelation(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist relation data", http.StatusInternalServerError)
		return
	}

	data := struct {
		models.Artist
		DatesLocations map[string][]string
	}{
		Artist:         artist,
		DatesLocations: relation.DatesLocations,
	}

	tmpl := template.Must(template.ParseFiles("./static/pages/artist.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
