package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	NotFoundHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "NotFoundHandler", 404, response.Code)
	}

	request, _ = http.NewRequest("GET", "/v1/token", nil)
	response = httptest.NewRecorder()

	NotFoundHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "NotFoundHandler", 404, response.Code)
	}
}

func TestTokenHandler(t *testing.T) {
	var token Token

	request, _ := http.NewRequest("POST", "/v1/token", nil)
	response := httptest.NewRecorder()

	TokenHandler(response, request)

	if response.Code != 200 {
		t.Errorf(errorMessage, "TokenHandler", 200, response.Code)
	}

	if err := token.Decode(response.Body.Bytes()); err != nil {
		t.Errorf(errorMessage, "TokenHandler", nil, err)
	}
	if len(token.Key) != 40 {
		t.Errorf(errorMessage, "TokenHandler", 40, len(token.Key))
	}
}
