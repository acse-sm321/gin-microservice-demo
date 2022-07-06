package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that GET request to home page return HTTP 200 and return the homepage

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	// define the route that use the same handler that the app is using
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		// strings.Index() was used to find first index of specified substr, return 0 if not found
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
