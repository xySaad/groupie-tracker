package handlers

import (
	"net/http"

	"groupie-tracker/config"
)

func Error(resp http.ResponseWriter, errMessage string, status int) {
	resp.WriteHeader(status)
	err := config.Templates.ExecuteTemplate(resp, "error.html", errMessage)
	if err != nil {
		resp.Write([]byte(errMessage))
	}
}
