package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testHandler = func(w http.ResponseWriter, r *http.Request) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func TestBasicAuth(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)

	os.Setenv("OFFERS_FINDER_PASSWORD", "secret")
	request.SetBasicAuth("locafox", "secret")

	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 200 {
		t.Errorf(errorMessage, "BasicAuth", 200, response.Code)
	}
}

func TestBasicAuthWrongMethod(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/token", nil)
	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 404 {
		t.Errorf(errorMessage, "BasicAuthWrongMethod", 404, response.Code)
	}
}

func TestBasicAuthNoAuthHeader(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 401 {
		t.Errorf(errorMessage, "BasicAuthNoAuthHeader", 401, response.Code)
	}
}

func TestBasicAuthWrongAuthHeader(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	request.Header.Add("Authorization", "a")

	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 401 {
		t.Errorf(errorMessage, "BasicAuthWrongAuthHeader", 401, response.Code)
	}
}

func TestBasicAuthWrongPayload(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	request.Header.Add("Authorization", "Basic a")

	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 401 {
		t.Errorf(errorMessage, "BasicAuthWrongPayload", 401, response.Code)
	}
}

func TestBasicAuthWrongCredentials(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/token", nil)
	request.SetBasicAuth("a", "a")

	response := httptest.NewRecorder()

	BasicAuth(testHandler(response, request))(response, request)

	if response.Code != 401 {
		t.Errorf(errorMessage, "BasicAuthWrongCredentials", 401, response.Code)
	}
}

func TestValidate(t *testing.T) {
	os.Setenv("OFFERS_FINDER_PASSWORD", "secret")

	if Validate("a", "b") {
		t.Errorf(errorMessage, "Validate", false, true)
	}

	if !Validate("locafox", "secret") {
		t.Errorf(errorMessage, "Validate", true, false)
	}
}
