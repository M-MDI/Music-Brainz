package main

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
)

type Artist struct {
    ID         int      `json:"id"`
    Name       string   `json:"name"`
    Members    []string `json:"members"`
    FirstAlbum string   `json:"firstAlbum"`
    StartYear  int      `json:"startYear"`
    Image      string   `json:"image"`
}

type Data struct {
    Artists []Artist
}

// Function to fetch artist data from the API
func fetchArtists() ([]Artist, error) {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var artists []Artist
    err = json.NewDecoder(resp.Body).Decode(&artists)
    if err != nil {
        return nil, err
    }
    return artists, nil
}

// Handler to serve the main page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    artists, err := fetchArtists()
    if err != nil {
        http.Error(w, "Failed to load artists", http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, Data{Artists: artists})
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", homeHandler)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
