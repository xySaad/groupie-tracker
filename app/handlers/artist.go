package handlers

import (
	"errors"
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/utils"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

type fetchError struct {
	status  int
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
	var bannerUrl string
	errChan := make(chan fetchError)

	go func() {
		err := utils.FetchData(BaseURL+"/artists/"+artistID, &artist)
		if err != nil {
			errChan <- fetchError{500, "Error fetching artist data", err}
			return
		}
		if artist["id"] == 0 {
			err = errors.New("404 - artist not found")
			errChan <- fetchError{404, err.Error(), err}
			return
		}
		bannerUrl, err = utils.GetBanner(artist["name"].(string))
		if err != nil {
			fmt.Println(err)
		}
		artist["banner"] = bannerUrl
		errChan <- fetchError{}
	}()

	go func() {
		err := utils.FetchData(BaseURL+"/relation/"+artistID, &relation)
		if err != nil {
			errChan <- fetchError{505, "Error fetching relation data", err}
			return
		}
		if artist["id"] == 0 {
			err = errors.New("artist not found")
			errChan <- fetchError{404, err.Error(), err}
			return
		}
		errChan <- fetchError{}
	}()

	for i := 0; i < 2; i++ {
		err := <-errChan
		if err.error != nil {
			fmt.Println(err)
			http.Error(w, err.message, err.status)
			return
		}
	}

	artist["datesLocations"] = relation["datesLocations"]

	err := config.Templates.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
