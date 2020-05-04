package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"strings"
	"testing"
)

func TestArticleUnauthenticated(t *testing.T) {
	r := GetRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHttpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		theArticle, _ := getArticleByID(1)
		title := theArticle.Title
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), fmt.Sprintf("<title>%s</title>", title)) > 0

		return statusOK && pageOK
	})
}

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

// Test that a GET request to the home page returns the list of articles
// in JSON format when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {
	r := GetRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHttpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to
		// an array of Article structs
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article
		err = json.Unmarshal(p, &articles)

		return err == nil && len(articles) >= 2 && statusOK
	})
}

// Test that a GET request to an article page returns the article in XML
// format when the Accept header is set to application/xml
func TestArticleXML(t *testing.T) {
	r := GetRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/article/view/:article_id", getArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHttpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to
		// an array of Article structs
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var a article
		err = xml.Unmarshal(p, &a)

		return err == nil && a.ID == 1 && len(a.Title) >= 0 && statusOK
	})
}