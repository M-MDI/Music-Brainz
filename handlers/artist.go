package handlers

import (
	api "Music-Brainz/api"
	config "Music-Brainz/config"
	funcs "Music-Brainz/funcs"
	"net/http"
	"strconv"
)


func Artist(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	data := config.MyStruct

	if config.ErrApiData != nil || config.ErrArtists != nil {
		ErrorHandler(w, "Network Authentication Required", http.StatusNetworkAuthenticationRequired)
		return
	}

	idTarget, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		ErrorHandler(w, "Error Bad Request!!", http.StatusBadRequest)
		return

	}

	if idTarget >= 1 && idTarget <= 52 {
		data.Artist = funcs.GetArtist(idTarget, data.Artists).(map[string]interface{})
		err := api.FetchData(data.Artist["concertDates"].(string), &data.Dates)

		if err != nil {
			ErrorHandler(w, "Internal Server Error !!", http.StatusInternalServerError)
			return
		}

		err = api.FetchData(data.Artist["locations"].(string), &data.Locations)

		if err != nil {
			ErrorHandler(w, "Internal Server Error !!", http.StatusInternalServerError)
			return
		}

		err = api.FetchData(data.Artist["relations"].(string), &data.Relation)

		if err != nil {
			ErrorHandler(w, "Internal Server Error !!", http.StatusInternalServerError)
			return
		}

		data.Dates["dates"] = funcs.ReValueDates(data.Dates["dates"].([]interface{}))
		err = funcs.ExctTmple(w, "artist.html", data)

		if err != nil {
			ErrorHandler(w, "Internal Server Error !!", http.StatusInternalServerError)
			return
		}

		} else {
		ErrorHandler(w, "Error Bad Request!!", http.StatusBadRequest)
		return
	}
}
