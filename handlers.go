package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func Base(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		log.Println(r.Method, r.URL, http.StatusNotFound)
		http.NotFound(w, r)
		return
	}
	log.Println(r.Method, r.URL, http.StatusOK)
	fmt.Fprintln(w, JsonResponse{"status": "OK"})
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, http.StatusOK)
	fmt.Fprintln(w, token)
}

func OffersHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("token") != token.Key {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, JsonResponse{"status": "unauthorized"})
		return
	}

	b, err := json.Marshal(offers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println(r.Method, r.URL, http.StatusOK)
	fmt.Fprintln(w, string(b))
}
