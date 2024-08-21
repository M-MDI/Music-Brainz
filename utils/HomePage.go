package Groupietracker

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := FetchArtists()
	if err != nil {
		http.Error(w, "Failed to load artists", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, Data{Artists: artists})
}
