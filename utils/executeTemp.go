package Groupietracker

import (
	"html/template"
	"net/http"
)

func ExctTmple(w http.ResponseWriter, name string, data any) error {
	var tpl *template.Template

	tpl, err := template.ParseGlob("/templates/*.html")
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}
