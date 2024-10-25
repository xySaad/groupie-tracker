package handlers

import (
	"net/http"
	"os"
)

func Home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "404 - page note found", http.StatusNotFound)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(res, "405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	homePage, err := os.ReadFile("./static/index.html")
	if err != nil {
		http.Error(res, "500 - internal server error", http.StatusInternalServerError)
		return
	}
	res.Write(homePage)
}
