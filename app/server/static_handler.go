package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func StaticHandler(res http.ResponseWriter, req *http.Request) {

	path := "./static"
	if req.URL.Path == "/" {
		path += "/index.html"
	} else {
		path += req.URL.Path
	}
	// Attempt to open the requested file
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		// Set the 404 status code
		res.WriteHeader(http.StatusNotFound)

		// Read the 404 page content and write it to the response
		page404, err := os.ReadFile("./static/pages/404.html")
		if err != nil {
			http.Error(res, "404 page not found", http.StatusNotFound)
			return
		}
		// Write the 404 page content to the response
		_, err = res.Write(page404)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if req.Method != "GET" {
		http.Error(res, "405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Add("Content-Length", strconv.Itoa(int(fileInfo.Size())))

	fileName := strings.Split(req.URL.Path, ".")

	switch fileName[len(fileName)-1] {
	case "js":
		res.Header().Add("Content-Type", "application/javascript")
	case "css":
		res.Header().Add("Content-Type", "text/css")
	case "html":
		res.Header().Add("Content-Type", "text/html")
	}
	file, err := os.ReadFile(path)

	if err != nil {
		http.Error(res, "500 - internal server error", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(200)

	_, err = res.Write(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error sending static file\n", err)
	}
}
