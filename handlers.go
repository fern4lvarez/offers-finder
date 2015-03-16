package main

import (
	"fmt"
	"log"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, "404 NOT FOUND")
	http.NotFound(w, r)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := token.Encode()
	if err != nil {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}

	log.Println(r.Method, r.URL, "200 OK")
	fmt.Fprintln(w, string(b))
}
