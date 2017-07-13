package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"net/url"
	"strings"
)

func TestStatus(t *testing.T) {

	uri := "http://acme.com/status"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	healthcheck_handler(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Received non-200 response: %d\n", resp.StatusCode)
	}
}

func TestRedirect(t *testing.T) {

	path := "/home/test"
	unlno := "997225821"
	param := make(url.Values)
	param["param1"] = []string{path}
	param["param2"] = []string{unlno}

	uri := "http://acme.com/some/url?" + param.Encode()

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	redirect_handler(w, req)
	resp := w.Result()

	if resp.StatusCode != 301 {
		t.Errorf("Received non-301 response: %d\n", resp.StatusCode)
	}
	if _, ok := resp.Header["Location"]; ok {
		if !strings.HasPrefix(resp.Header["Location"][0], "https") {
			t.Errorf("Received non https Location header: %s\n", resp.Header["Location"][0])
		}
		if resp.Header["Location"][0][5:] != uri[4:] {
			t.Errorf("Location header %s does not match: %s\n", resp.Header["Location"][0], uri)
		}
	} else {
		t.Error("Location header missing from response\n")
	}
}
