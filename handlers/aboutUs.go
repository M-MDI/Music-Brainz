package handlers

import (
	funcs "Music-Brainz/funcs"
	"net/http"
)


func AboutUs(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	funcs.ExctTmple(w, "about.html", nil)
}