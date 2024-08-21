package Groupietracker

import (
	"encoding/json"
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
func FetchArtists() ([]Artist, error) {
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
