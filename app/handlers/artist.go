package handlers

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/utils"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

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

	msg, status := utils.ValidateArtistID(queries)

	if status != http.StatusOK {
		http.Error(w, msg, status)
		return
	}

	artistID := queries["id"][0]

	var artist utils.Object
	var relation utils.Object
	errChan := make(chan error)

	go func() {
		err := utils.FetchData(BaseURL+"/artists/"+artistID, &artist)
		errChan <- err
	}()

	go func() {
		err := utils.FetchData(BaseURL+"/relation/"+artistID, &relation)
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		if <-errChan != nil {
			http.Error(w, "Error fetching artist data from relation", http.StatusInternalServerError)
			return
		}
	}

	var datesLocations utils.Object
	err := relation.Get(&datesLocations, ".datesLocations")
	if err != nil {
		http.Error(w, "Error getting datesLocations from relation", http.StatusInternalServerError)
		return
	}

	artist["datesLocations"] = datesLocations

	err = config.Templates.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
