package main

import (
	"log"
	"net/http"
)

var token = NewToken()

func main() {
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/v1/token", BasicAuth(TokenHandler))

	log.Println("Listening and serving HTTP on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
