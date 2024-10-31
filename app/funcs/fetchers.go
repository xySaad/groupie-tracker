package funcs

import (
	"encoding/json"
	"fmt"
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

func FetchArtistRelation(artistID int) (ArtistRelation, error) {
	var relation ArtistRelation
	err := FetchData(fmt.Sprintf("%s/relation/%d", BaseURL, artistID), &relation)
	return relation, err
}

// Fetch artist dates
func FetchArtistDates(artistID int) (ArtistDates, error) {
	var dates ArtistDates
	err := FetchData(fmt.Sprintf("%s/dates/%d", BaseURL, artistID), &dates)
	for i := range dates.Dates {
		dates.Dates[i] = dates.Dates[i][1:]
	}
	return dates, err
}

// Fetch all artists
func FetchAllArtists() ([]Artist, error) {
	var artists []Artist
	err := FetchData(fmt.Sprintf("%s/artists", BaseURL), &artists)
	return artists, err
}

// Fetch specific artist data
func FetchArtist(artistID int) (Artist, error) {
	var artist Artist
	err := FetchData(fmt.Sprintf("%s/artists/%d", BaseURL, artistID), &artist)
	return artist, err
}
