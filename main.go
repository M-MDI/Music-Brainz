package main

import (
	Gr "Groupietracker/utils"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Gr.HomeHandler)
	http.HandleFunc("/artist/{id}", Gr.ArtistDetails)
	log.Println("Server starting on : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
