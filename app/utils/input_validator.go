package utils

import (
	"net/http"
	"strconv"
)

func ValidateArtistID(queries map[string][]string) (string, int) {
	if queries["id"] == nil {
		return "400 - missing id", http.StatusBadRequest
	}

	artistIDStr := queries["id"][0]

	if artistIDStr == "" {
		return "400 - empty artist id", http.StatusBadRequest
	}

	artistID, err := strconv.Atoi(artistIDStr)

	if err != nil {
		return "400 - invalid artist id", http.StatusBadRequest
	}

	if artistID < 1 || artistID > 52 {
		return "404 - artist not found", http.StatusNotFound
	}

	return "ok", http.StatusOK
}
