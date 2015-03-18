package main

import (
	"log"
	"net/http"
)

var (
	token  = NewToken()
	offers = []Offer{NewOffer(0),
		NewOffer(1),
		NewOffer(2),
		NewOffer(3),
		NewOffer(4),
		NewOffer(5),
	}
)

func main() {
	http.HandleFunc("/", Base(IndexHandler))
	http.HandleFunc("/v1/token", Base(Post(BasicAuth(TokenHandler))))
	http.HandleFunc("/v1/offers", Base(Post(BasicAuth(OffersHandler))))

	log.Println("Listening and serving HTTP on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
