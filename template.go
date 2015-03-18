package main

import (
	"net/http"
)

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
