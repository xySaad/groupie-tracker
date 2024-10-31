package utils

import (
	"net/http"
	"strconv"
)

func ValidateArtistID(queries map[string][]string) (int, string, int) {
	msg := "ok"
	status := http.StatusOK

	if queries["id"] == nil {
		return 0, "400 - missing id", http.StatusBadRequest
	}

	artistIDStr := queries["id"][0]

	artistID, err := strconv.Atoi(artistIDStr)

	if artistID < 1 || artistID > 52 {
		msg, status = "404 - artist not found", http.StatusNotFound
	}

	if err != nil {
		msg, status = "400 - invalid artist id", http.StatusBadRequest
	}

	if artistIDStr == "" {
		msg, status = "400 - invalid empty artist id", http.StatusBadRequest
	}

	return artistID, msg, status
}
