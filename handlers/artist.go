package handlers

import (
	api "Music-Brainz/api"
	config "Music-Brainz/config"
	funcs "Music-Brainz/funcs"
	"net/http"
	"strconv"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	data := config.MyStruct
	if config.ErrApiData != nil || config.ErrArtists != nil {
		http.Error(w, http.StatusText(http.StatusNetworkAuthenticationRequired), http.StatusNetworkAuthenticationRequired)
		return
	}
	idTarget, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if idTarget >= 1 && idTarget <= 52 {

		data.Artist = funcs.GetArtist(idTarget, data.Artists).(map[string]interface{})
		err := api.FetchData(data.Artist["concertDates"].(string), &data.Dates)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = api.FetchData(data.Artist["locations"].(string), &data.Locations)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = api.FetchData(data.Artist["relations"].(string), &data.Relation)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		data.Dates["dates"] = funcs.ReValueDates(data.Dates["dates"].([]interface{}))
		err = funcs.ExctTmple(w, "artist.html", data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
