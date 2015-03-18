package main

import (
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

func init() {
	Token_ = NewToken()
	Offers = []Offer{NewOffer(0),
		NewOffer(1),
		NewOffer(2),
		NewOffer(3),
		NewOffer(4),
		NewOffer(5),
	}
	Username = "locafox"
	Password = os.Getenv("OFFERS_FINDER_PASSWORD")
	if Password == "" {
		Password = "LocaF#xes!"
	}
}

func main() {
	http.HandleFunc("/", Base(IndexHandler))
	http.HandleFunc("/v1/token", Base(Post(BasicAuth(TokenHandler))))
	http.HandleFunc("/v1/offers", Base(Post(BasicAuth(OffersHandler))))

	log.Println("Listening and serving HTTP on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
