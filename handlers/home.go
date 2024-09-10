package handlers

import (
	// api "Music-Brainz/static/api"

	config "Music-Brainz/config"
	"Music-Brainz/funcs"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		err := config.ErrApiData
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNetworkAuthenticationRequired), http.StatusNetworkAuthenticationRequired)
			return
		}

		data := config.MyStruct
		err = funcs.ExctTmple(w, "index.html", data)
		if err != nil {
			http.Error(w, "Failed to render template: ", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
