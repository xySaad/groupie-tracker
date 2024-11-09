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

func FetchArtistLocations(artistID int) (models.ArtistLocations, error) {
	var locations models.ArtistLocations
	err := FetchData(fmt.Sprintf("%s/locations/%d", models.BaseURL, artistID), &locations)
	return locations, err
}

func FetchArtistRelation(artistID int) (models.ArtistRelation, error) {
	var relation models.ArtistRelation
	err := FetchData(fmt.Sprintf("%s/relation/%d", models.BaseURL, artistID), &relation)
	return relation, err
}

// Fetch artist dates
func FetchArtistDates(artistID int) (models.ArtistDates, error) {
	var dates models.ArtistDates
	err := FetchData(fmt.Sprintf("%s/dates/%d", models.BaseURL, artistID), &dates)
	for i := range dates.Dates {
		dates.Dates[i] = dates.Dates[i][1:]
	}
	return dates, err
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
