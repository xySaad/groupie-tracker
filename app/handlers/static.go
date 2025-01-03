package handlers

import (
	"net/http"
	"os"
)

func Static(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(resp, "405 - method not allowed", 405)
		return
	}
	fileInfo, err := os.Stat(req.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			Error(resp, "404 - page not found", 404)
			return
		}
		Error(resp, "500 - internal server error", 500)
		return
	}
	if fileInfo.IsDir() {
		Error(resp, "404 - page not found", 404)
		return
	}
	http.ServeFile(resp, req, req.URL.Path[1:])
}
