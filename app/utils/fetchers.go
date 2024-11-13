package utils

import (
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
)

// Utility function for fetching data
func FetchData(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	getter, err := Decode(string(body))
	if err != nil {
		return err
	}
	return getter.Get(v, "")
}

func FetchArtistRelation(artistID int) (Object, error) {
	var relation Object
	err := FetchData(fmt.Sprintf("%s/relation/%d", models.BaseURL, artistID), &relation)
	return relation, err
}

// Fetch all artists
func FetchAllArtists() ([]Object, error) {
	var artists []Object
	err := FetchData(fmt.Sprintf("%s/artists", models.BaseURL), &artists)
	return artists, err
}

// Fetch specific artist data
func FetchArtist(artistID int) (Object, error) {
	var artist Object
	err := FetchData(fmt.Sprintf("%s/artists/%d", models.BaseURL, artistID), &artist)
	return artist, err
}
