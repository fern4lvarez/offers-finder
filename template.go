package main

import (
	"net/http"
)

// RenderMapTemplate renders the map.html template
// with existing Offers
func RenderMapTemplate(w http.ResponseWriter) {
	data := struct {
		Offers []Offer
	}{
		Offers: Offers,
	}

	err := MapTemplate.ExecuteTemplate(w, "map.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
