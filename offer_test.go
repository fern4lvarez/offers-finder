package main

import (
	"bytes"
	"testing"
)

func TestNewOffer(t *testing.T) {
	offer := NewOffer(0)

	if offer.Id != 0 {
		t.Errorf(errorMessage, "NewOffer", 0, offer.Id)
	}

	if offer.Lat < 52.47 || offer.Lat > 52.55 {
		t.Errorf(errorMessage, "NewOffer",
			"52.47<offer.Lat<52.55", offer.Lat)
	}

	if offer.Long < 13.24 || offer.Long > 13.50 {
		t.Errorf(errorMessage, "NewOffer",
			"13.24<offer.Long<13.50", offer.Long)
	}
}

func TestOfferEncode(t *testing.T) {
	offer := Offer{
		Id:   3,
		Lat:  52.51,
		Long: 13.38,
	}
	expectedEncoded := []byte(`{"id":3,"lat":52.51,"long":13.38}`)

	b, err := offer.Encode()

	if err != nil {
		t.Errorf(errorMessage, "OfferEncode", nil, err)
	}
	if bytes.Compare(b, expectedEncoded) != 0 {
		t.Errorf(errorMessage, "OfferEncode", string(b), string(expectedEncoded))
	}
}

func TestGenerateCoordinates(t *testing.T) {
	for i := 0; i < 5; i++ {
		lat, long := generateCoordinates()
		if lat < 52.47 || lat > 52.55 {
			t.Errorf(errorMessage, "generateCoordinates",
				"52.47<lat<52.55", lat)
		}

		if long < 13.24 || long > 13.50 {
			t.Errorf(errorMessage, "generateCoordinates",
				"13.24<long<13.50", long)
		}
	}
}

func TestRandomCoordinate(t *testing.T) {
	c1 := randomCoordinate(10, 11)
	c2 := randomCoordinate(10, 11)
	c3 := randomCoordinate(10, 11)
	if c1 < 10 || c1 > 11 {
		t.Errorf(errorMessage, "randomCoordinate", "10<c1<11", c1)
	}

	if c2 < 10 || c2 > 11 {
		t.Errorf(errorMessage, "randomCoordinate", "10<c2<11", c2)
	}

	if c3 < 10 || c3 > 11 {
		t.Errorf(errorMessage, "randomCoordinate", "10<c2<11", c2)
	}

	if c1 == c2 && c2 == c3 && c3 == c1 {
		t.Errorf(errorMessage, "randomCoordinate", "c1!=c2!=c3", c1)
	}

}
