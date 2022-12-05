package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// Status code test
	if w.Code != 200 {
		t.Error("Response status code is not equal to 200")
	}

	// Return value test
	if w.Body.String() != "pong" {
		t.Error("Response body is not equal to pong")
	}
}
