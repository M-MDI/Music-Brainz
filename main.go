package main

import (
	Gr "Groupietracker/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Gr.HomeHandler)
	http.HandleFunc("/artist/{id}", Gr.ArtistDetails)
	log.Println("Server starting on : http://localhost:8100")
	if err := http.ListenAndServe(":8100", nil); err != nil {
		fmt.Println(err)
		return
	}
}
