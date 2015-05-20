package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Token_ is a unique 40 characters token
var Token_ *Token

// Offers is a random list of offers based in Berlin
var Offers []Offer

// Username and Password are the server's Basic Auth credentials
var Username, Password string

// DisplayEndpoint is a unique key that is used as a secret
// endpoint to the display of the offers map
var DisplayEndpoint string

// MapTemplate of the offers finder map
var MapTemplate *template.Template

func init() {
	Token_ = NewToken()
	Offers = []Offer{NewOffer(0),
		NewOffer(1),
		NewOffer(2),
		NewOffer(3),
		NewOffer(4),
		NewOffer(5),
	}
	Username = "user"
	Password = os.Getenv("OFFERS_FINDER_PASSWORD")
	if Password == "" {
		Password = "secret"
	}

	DisplayEndpoint = fmt.Sprintf("/%s", generateToken(16))

	MapTemplate = template.Must(template.ParseFiles("templates/map.html"))
}

func main() {
	http.HandleFunc("/", Base(IndexHandler))
	http.HandleFunc(DisplayEndpoint, DisplayHandler)
	http.HandleFunc("/v1/token", Base(Post(BasicAuth(TokenHandler))))
	http.HandleFunc("/v1/offers", Base(Post(BasicAuth(OffersHandler))))
	http.HandleFunc("/v1/offers/display", Base(Post(BasicAuth(OffersHandler))))

	log.Println("Listening and serving HTTP on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
