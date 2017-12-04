package main

import (
	"fmt"
	"net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheck(t *testing.T) {

	fmt.Println("Checking health check")
	req, _ := http.NewRequest("GET", "/health", nil)
	res := httptest.NewRecorder()

	HealthCheck(res, req)

	if res.Body.String() != "Hey I am working here /health!" {
		t.Error("Fail! It should return: Hey I am working here /health!")
	}
}

func TestRootHandler(t *testing.T) {
	
		fmt.Println("Checking main handler")
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
	
		Hanlder(res, req)
		if res.Body.String() != "Hello, /" {
			t.Error("Fail! It should return: Hello, /")
		}
}