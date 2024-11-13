package handlers

import (
	"fmt"
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
	var datesLocations utils.Object
	err = relation.Get(&datesLocations, ".datesLocations")
	if err != nil {
		http.Error(w, "Error getting datesLocations from relation", http.StatusInternalServerError)
		return
	}
	data := struct {
		Artist         utils.Object
		DatesLocations utils.Object
	}{
		Artist:         artist,
		DatesLocations: datesLocations,
	}

	tmpl, err := template.ParseFiles("./static/pages/artist.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
