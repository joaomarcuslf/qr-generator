package handlers

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title       string
	Description string
}

func Index(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title:       "QR Code Generator",
		Description: "A page to generate QR",
	}

	t, err := template.ParseFiles("templates/generator.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, p)
}
