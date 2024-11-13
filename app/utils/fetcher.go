package utils

import (
	"io"
	"net/http"
)

func FetchData[Type *Object | *[]Object](url string, v Type) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	getter, err := Decode(string(body))
	if err != nil {
		return err
	}
	return getter.Get(v, "")
}
