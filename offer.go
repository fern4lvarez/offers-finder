package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

// Offer containing its location in coordinates
type Offer struct {
	Id   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// NewOffer with a randomly generated coordinates
func NewOffer(id int) Offer {
	lat, long := generateCoordinates()

	return Offer{
		Id:   id,
		Lat:  lat,
		Long: long,
	}
}

// Encode encodes an Offer into json in bytes
func (offer Offer) Encode() (b []byte, err error) {
	b, err = json.Marshal(offer)
	return
}

// generateCoordinates retuns random coordinates based
// in Berlin
func generateCoordinates() (float64, float64) {
	lat := randomCoordinate(52.47, 52.55)
	long := randomCoordinate(13.24, 13.50)
	return lat, long
}

// randomCoordinate returns a coordinate given
// a range
func randomCoordinate(min, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Float64()*(max-min) + min
}
