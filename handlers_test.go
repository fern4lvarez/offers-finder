package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBase(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	response := httptest.NewRecorder()

	Base(testHandler(response, request))(response, request)

	if aclo := response.Header().Get("Access-Control-Allow-Origin"); aclo != "*" {
		t.Errorf(errorMessage, "Base", "*", aclo)
	}

	if aclh := response.Header().Get("Access-Control-Allow-Headers"); aclh !=
		"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization" {
		t.Errorf(errorMessage, "Base",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			aclh)
	}

	if ct := response.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf(errorMessage, "Base", "application/json", ct)
	}
}

func TestPostOK(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	response := httptest.NewRecorder()

	Post(testHandler(response, request))(response, request)

	if response.Code != 200 {
		t.Errorf(errorMessage, "PostOK", 200, response.Code)
	}
}

func TestPostNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/token", nil)
	response := httptest.NewRecorder()

	Post(testHandler(response, request))(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "PostUnauthorized", 404, response.Code)
	}
}

func TestIndexHandlerOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	if response.Code != 200 {
		t.Errorf(errorMessage, "IndexHandlerOK", 200, response.Code)
	}
}

func TestIndexHandlerNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/foo", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "IndexHandlerNotFound", 404, response.Code)
	}
}

func TestIndexHandlerPost(t *testing.T) {
	request, _ := http.NewRequest("POST", "/", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "IndexHandlerPost", 404, response.Code)
	}
}

func TestDisplayHandlerOK(t *testing.T) {
	request, _ := http.NewRequest("GET", DisplayEndpoint, nil)
	response := httptest.NewRecorder()

	DisplayHandler(response, request)

	if response.Code != 200 {
		t.Errorf(errorMessage, "DisplayHandlerOK", 200, response.Code)
	}
}

func TestDisplayHandlerNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/foo", nil)
	response := httptest.NewRecorder()

	DisplayHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "DisplayHandlerNotFound", 404, response.Code)
	}
}

func TestDisplayHandlerPost(t *testing.T) {
	request, _ := http.NewRequest("POST", DisplayEndpoint, nil)
	response := httptest.NewRecorder()

	DisplayHandler(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "DisplayHandlerPost", 404, response.Code)
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
