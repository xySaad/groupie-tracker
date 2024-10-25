package server

import (
	"groupie-tracker/handlers"
	"net/http"
)

func router(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		handlers.Home(res, req)
	default:
		handlers.Static(res, req)
	}
}
