package main

import (
	"log"
	"net/http"
	Gr "Groupietracker/utils"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Gr.HomeHandler)
	log.Println("Server starting on : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
