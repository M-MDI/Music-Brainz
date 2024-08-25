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
		if r.URL.Path != "/" {
			tmpl, err := template.ParseFiles("templates/errors.html")
				err = tmpl.Execute(w, err)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}
		return
	}

			tmpl := template.Must(template.ParseFiles("templates/HomePage.html"))
			tmpl.Execute(w, Data{Artists: artists})

}