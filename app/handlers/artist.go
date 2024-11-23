package handlers

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/utils"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

type fetchError struct {
	message string
	error
}

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
	errChan := make(chan fetchError)

	go func() {
		err := utils.FetchData(BaseURL+"/artists/"+artistID, &artist)
		if err != nil {
			errChan <- fetchError{"Error fetching artist data", err}
		} else {
			errChan <- fetchError{}
		}
	}()

	go func() {
		err := utils.FetchData(BaseURL+"/relation/"+artistID, &relation)
		if err != nil {
			errChan <- fetchError{"Error fetching relation data", err}
		} else {
			errChan <- fetchError{}
		}
	}()

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err.error != nil {
			http.Error(w, err.message, http.StatusInternalServerError)
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
