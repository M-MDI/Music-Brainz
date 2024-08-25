package Groupietracker

import (
	"html/template"
	"net/http"
	"strconv"
)

func ArtistDetails(w http.ResponseWriter, r *http.Request) {

	Deta.Art, _ = FetchArtists()

	// Handle errors for each Fetch call
	if err := Fetch("https://groupietrackers.herokuapp.com/api/dates", &Deta.Dates); err != nil {
		http.Error(w, "Failed to fetch dates", http.StatusInternalServerError)
		return
	}
	if err := Fetch("https://groupietrackers.herokuapp.com/api/relation", &Deta.Relations); err != nil {
		http.Error(w, "Failed to fetch relations", http.StatusInternalServerError)
		return
	}
	if err := Fetch("https://groupietrackers.herokuapp.com/api/locations", &Deta.Locations); err != nil {
		http.Error(w, "Failed to fetch locations", http.StatusInternalServerError)
		return
	}
	ID := r.PathValue("id")
	id, err := strconv.Atoi(ID)

	if err != nil {
		tmpl, err := template.ParseFiles("templates/errors.html")
		err = tmpl.Execute(w, err)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
		return
	}

	if id > 0 && id <= len(Deta.Art) {
		tmpl, err := template.ParseFiles("templates/Art.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		response := Resp{
			Art:       Deta.Art[id-1],
			Dates:     Deta.Dates.Index[id-1].Dates,
			Locations: Deta.Locations.Index[id-1].Location,
			Relations: Deta.Relations.Index[id-1].Relation,
		}
		err = tmpl.Execute(w, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		tmpl, err := template.ParseFiles("templates/errors.html")
				err = tmpl.Execute(w, err)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}
	}
}
