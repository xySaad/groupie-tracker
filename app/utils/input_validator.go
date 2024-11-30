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

	id, err := strconv.Atoi(artistIDStr)

	if err != nil {
		return "400 - invalid artist id", http.StatusBadRequest
	}

	if id <= 0 {
		return "404 - artist not found", http.StatusNotFound
	}
	return "ok", http.StatusOK
}
