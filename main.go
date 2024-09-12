package main

import (
	"fmt"
	handlers "Music-Brainz/handlers"
	"net/http"
)

func main() {
	port := "8082"
	fmt.Println("Server started at http://localhost:" + port)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/src/", handlers.CssHandler)
	http.HandleFunc("/artist", handlers.Artist)
	http.HandleFunc("/about-us", handlers.AboutUs)
	http.ListenAndServe(":"+port, nil)
}
