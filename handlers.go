package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Handler acts as a type of a basic handler
type Handler func(w http.ResponseWriter, r *http.Request)

// Base wrapper that sets basic headers
func Base(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

// Post wrapper only accepts POST requests
func Post(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only POST requests are authorized
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}
}

// IndexHandler handles GET requests to the root url
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		log.Println(r.Method, r.URL, http.StatusNotFound)
		http.NotFound(w, r)
		return
	}
	log.Println(r.Method, r.URL, http.StatusOK)
	fmt.Fprintln(w, JsonResponse{"status": "OK"})
}

// DisplayHandler handles GET requests to the display map
// secret endpoint
func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != DisplayEndpoint {
		log.Println(r.Method, r.URL, http.StatusNotFound)
		http.NotFound(w, r)
		return
	}
	log.Println(r.Method, r.URL, http.StatusOK)
	RenderMapTemplate(w)
}

// TokenHandler handles token requests
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, http.StatusOK)
	fmt.Fprintln(w, Token_)
}

// OffersHanlder handles offers requests, only accepts requests
// with the right payload
func OffersHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("token") != Token_.Key {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, JsonResponse{"status": "unauthorized"})
		return
	}

	if strings.HasSuffix(r.URL.Path, "display") {
		log.Println(r.Method, r.URL, http.StatusMovedPermanently)
		http.Redirect(w, r, "http://"+r.Host+DisplayEndpoint, 301)
		return
	}

	log.Println(r.Method, r.URL, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(Offers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, string(b))
}
