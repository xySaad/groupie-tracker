package main

import (
	"io"
	"net/http"
	"strings"
)

type responseStruct struct {
	status int
	body   string
}

func containAll(s string, items []string) (bool, string) {
	for _, item := range items {
		if !strings.Contains(s, item) {
			return false, item
		}
	}
	return true, ""
}

func request(method, path string) (responseStruct, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, "http://0.0.0.0:8080"+path, nil)
	if err != nil {
		return responseStruct{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return responseStruct{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseStruct{}, err
	}

	return responseStruct{resp.StatusCode, string(body)}, nil
}
