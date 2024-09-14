package handlers

import (
	"net/http"
	"strconv"

	"Music-Brainz/funcs"
)

type stError struct {
	Code     []string
	MsgError string
}

func ConvSlice(str string) []string {

	var slice []string

	for _, v := range str {

		slice =append(slice, string(v))
	}

	return slice
}

func ErrorHandler(w http.ResponseWriter, ErrorMsg string, StatusCode int) {
	
	w.WriteHeader(StatusCode)

	var Err  stError

	Err.Code = ConvSlice(strconv.Itoa(StatusCode))
	Err.MsgError = ErrorMsg

	tempErr := funcs.ExctTmple(w, "error.html", Err)
	if tempErr != nil {
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
