package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"strings"
	"testing"
)

// Test that GET request to the Home page returns home page with HTTP Code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := GetRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHttpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}