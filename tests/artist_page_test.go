package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestQueenMembers(t *testing.T) {
	resp, err := request(http.MethodGet, "/artist?id=1")
	if err != nil {
		t.Error(err)
	}
	if resp.status != 200 {
		t.Error("expected status code 200", "got:", resp.status)
	}

	expected := []string{
		"Freddie Mercury",
		"Brian May",
		"John Daecon",
		"Roger Meddows-Taylor",
		"Mike Grose",
		"Barry Mitchell",
		"Doug Fogie",
	}
	valid, missing := containAll(resp.body, expected)
	if !valid {
		t.Error("missing:", missing)
	}
}

func TestGorillazFirstAlbum(t *testing.T) {
	resp, err := request(http.MethodGet, "/artist?id=39")
	if err != nil {
		t.Error(err)
	}
	if resp.status != 200 {
		t.Error("expected status code 200", "got:", resp.status)
	}
	if !strings.Contains(resp.body, "26-03-2001") {
		t.Error("missing 26-03-2001")
	}
}

func TestTravisScottLocations(t *testing.T) {
	resp, err := request(http.MethodGet, "/artist?id=30")
	if err != nil {
		t.Error(err)
	}
	if resp.status != 200 {
		t.Error("expected status code 200", "got:", resp.status)
	}

	expected := []string{
		"Atlanta, Usa",
		"Frauenfeld, Switzerland",
		"Houston, Usa",
		"London, Uk",
		"Los Angeles, Usa",
		"New Orleans, Usa",
		"Philadelphia, Usa",
		"Santiago, Chile",
		"Sao Paulo, Brazil",
		"Turku, Finland",
	}

	valid, missing := containAll(resp.body, expected)
	if !valid {
		t.Error("missing:", missing)
	}
}

func TestFooFightersMembers(t *testing.T) {
	resp, err := request(http.MethodGet, "/artist?id=51")
	if err != nil {
		t.Error(err)
	}
	if resp.status != 200 {
		t.Error("expected status code 200", "got:", resp.status)
	}

	expected := []string{
		"Dave Grohl",
		"Nate Mendel",
		"Taylor Hawkins",
		"Chris Shiflett",
		"Pat Smear",
		"Rami Jaffee",
	}
	valid, missing := containAll(resp.body, expected)
	if !valid {
		t.Error("missing:", missing)
	}
}
