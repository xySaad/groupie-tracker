package funcs

import (
	"groupie-tracker/utils"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 - page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "501 - method not implemented", http.StatusNotImplemented)
		return
	}

	artists, err := FetchAllArtists()
	if err != nil {
		HandleError(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err = tmpl.Execute(w, artists)
	if err != nil {
		HandleError(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

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
		HandleError(w, msg, status)
		return
	}

	artist, err := FetchArtist(artistID)
	if err != nil {
		HandleError(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	relation, err := FetchArtistRelation(artistID)
	if err != nil {
		HandleError(w, "Error fetching artist relation data", http.StatusInternalServerError)
		return
	}

	dates, err := FetchArtistDates(artistID)
	if err != nil {
		HandleError(w, "Error fetching artist dates", http.StatusInternalServerError)
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
		HandleError(w, "Error rendering template", http.StatusInternalServerError)
	}
}
