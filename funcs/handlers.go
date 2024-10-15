package funcs

import (
	"net/http"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	artists, err := FetchAllArtists()
	if err != nil {
		HandleError(w, err, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err = tmpl.Execute(w, artists)
	if err != nil {
		HandleError(w, err, "Error rendering template", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/artist" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	artistIDStr := r.URL.Query().Get("id")
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil || artistID < 1 || artistID > 52 {
		HandleError(w, err, "Invalid artist ID", http.StatusNotFound)
		return
	}

	artist, err := FetchArtist(artistID)
	if err != nil {
		HandleError(w, err, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	relation, err := FetchArtistRelation(artistID)
	if err != nil {
		HandleError(w, err, "Error fetching artist relation data", http.StatusInternalServerError)
		return
	}

	dates, err := FetchArtistDates(artistID)
	if err != nil {
		HandleError(w, err, "Error fetching artist dates", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artist
		Dates          []string
		DatesLocations map[string][]string
	}{
		Artist:         artist,
		Dates:          dates.Dates,
		DatesLocations: relation.DatesLocations,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w, err, "Error rendering template", http.StatusInternalServerError)
	}
}
