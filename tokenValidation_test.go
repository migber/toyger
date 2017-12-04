package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestTokenValidationEvent(t *testing.T) {

	req, _ := http.NewRequest("GET", "/events", nil)
	req.Header.Add("authorization", "Bearer ")	
	res := httptest.NewRecorder()

	TeamsHandler(res, req)

	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Error("Should status code be unauthorized")
	}
}

func TestTokenValidationCommissaire(t *testing.T) {
	
	req, _ := http.NewRequest("GET", "/commissaires", nil)
	req.Header.Add("authorization", "Bearer ")	
	res := httptest.NewRecorder()

	CommissairesHandler(res, req)

	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Error("Should status code be unauthorized")
	}
}

func TestTokenValidationTeams(t *testing.T) {
	
	req, _ := http.NewRequest("GET", "/teams", nil)
	req.Header.Add("authorization", "Bearer ")	
	res := httptest.NewRecorder()

	TeamsHandler(res, req)

	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Error("Should status code be unauthorized")
	}
}