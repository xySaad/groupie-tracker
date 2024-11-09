package utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

// Utility function for fetching data
func FetchData(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

func FetchArtistRelation(artistID int) (models.ArtistRelation, error) {
	var relation models.ArtistRelation
	err := FetchData(fmt.Sprintf("%s/relation/%d", models.BaseURL, artistID), &relation)
	return relation, err
}

// Fetch all artists
func FetchAllArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := FetchData(fmt.Sprintf("%s/artists", models.BaseURL), &artists)
	return artists, err
}

// Fetch specific artist data
func FetchArtist(artistID int) (models.Artist, error) {
	var artist models.Artist
	err := FetchData(fmt.Sprintf("%s/artists/%d", models.BaseURL, artistID), &artist)
	return artist, err
}
