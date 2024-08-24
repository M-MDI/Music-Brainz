package Groupietracker

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		Id       int                 `json:"id"`
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Location struct {
	Index []struct {
		Id       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

type Details struct {
	Art       []Artist
	Dates     Date
	Locations Location
	Relations Relation
}

type Resp struct {
	Art       Artist
	Dates     []string
	Locations []string
	Relations map[string][]string
}

type Data struct {
	Artists []Artist
}

var Deta Details


var artists []Artist