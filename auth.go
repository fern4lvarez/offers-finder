package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

var authMessage = "Authenticating %s %s request... %s"

// BasicAuth acts as a wrapper of the handler, providing basic auth
func BasicAuth(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only POST requests are authorized
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		// Request with no Authorization Header
		authHeader, ok := r.Header["Authorization"]
		if !ok {
			log.Printf(authMessage, r.Method, r.URL, "unauthorized")
			w.Header().Set("Www-Authenticate", "Basic realm=\"Authorization Required\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Authorization Header with wrong format
		auth := strings.SplitN(authHeader[0], " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			log.Printf(authMessage, r.Method, r.URL, "unauthorized")
			w.Header().Set("Www-Authenticate", "Basic realm=\"Authorization Required\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Authorization payload wrong encoded
		payload, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			log.Printf(authMessage, r.Method, r.URL, "unauthorized")
			w.Header().Set("Www-Authenticate", "Basic realm=\"Authorization Required\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Wrong Authorization credentials
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			log.Printf(authMessage, r.Method, r.URL, "unauthorized")
			w.Header().Set("Www-Authenticate", "Basic realm=\"Authorization Required\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Printf(authMessage, r.Method, r.URL, "OK")
		handler(w, r)
	}
}

// Validate username and password from basic auth
func Validate(username, password string) bool {
	return username == "locafox" && password == "LocaF#xes!"
}
