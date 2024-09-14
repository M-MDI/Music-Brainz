package handlers

import (
	config "Music-Brainz/config"
	funcs "Music-Brainz/funcs"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	
	if r.URL.Path == "/" {
	
		err := config.ErrApiData
	
		if err != nil {
			ErrorHandler(w, "Network Authentication Required", http.StatusNetworkAuthenticationRequired)
			return
		}
	
		data := config.MyStruct
		err = funcs.ExctTmple(w, "index.html", data)
	
		if err != nil {
			ErrorHandler(w, "Failed to render template: ", http.StatusInternalServerError)
			return
		}
	
		} else {
		ErrorHandler(w, "This Page Not Found !!", http.StatusNotFound)
		return
	}
}
