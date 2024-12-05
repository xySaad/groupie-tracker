package main

import (
	"net/http"
	"testing"
)

var Methods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

type pathStruct struct {
	path   string
	method string
	query  string
}

func (p *pathStruct) Get(subPath string) string {
	return p.path + subPath + "?" + p.query
}

var paths = [...]pathStruct{
	{
		path:   "/",
		method: http.MethodGet,
	},
	{
		path:   "/artist",
		method: http.MethodGet,
		query:  "id=1",
	},
}

func TestMethodsAndInvalidPath(t *testing.T) {
	for _, method := range Methods {
		for _, path := range paths {
			resp, err := request(method, path.Get(""))
			if err != nil {
				t.Error(err)
			}
			if method == path.method {
				if resp.status != http.StatusOK {
					t.Error("expected status code 200", "got:", resp.status)
				}
			} else if resp.status != http.StatusMethodNotAllowed {
				t.Error("expected status code 405", "got:", resp.status)

			}

			// test error priority
			// 404 should be prioritiezed over invalid 405
			resp, err = request(method, path.Get("/invalidpath"))
			if err != nil {
				t.Error(err)
			}
			if resp.status != http.StatusNotFound {
				t.Error("expected status code 404", "got:", resp.status)
			}
		}
	}
}
