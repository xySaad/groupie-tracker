package models

const BaseURL = "https://groupietrackers.herokuapp.com/api"

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertDates string   `json:"concertDates"`
}

type ArtistDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ArtistLocations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type ArtistRelation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
