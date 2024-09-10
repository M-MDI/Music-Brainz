package handlers

import (
	funcs "Music-Brainz/funcs"
	"net/http"
)

func AboutUs(w http.ResponseWriter, r *http.Request) {
	funcs.ExctTmple(w, "about.html", nil)
}
