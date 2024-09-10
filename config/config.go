package config

import (
	api "Music-Brainz/api"
)

var (
	// The Data Who Send To The Templates
	MyStruct struct {
		Artists   []map[string]interface{}
		Artist    map[string]interface{}
		Dates     map[string]interface{}
		Locations map[string]interface{}
		Relation  map[string]interface{}
	}

	// For Stor The global Api
	ApiData map[string]interface{}

	// The Url Of Global Api
	ApiUrl = "https://groupietrackers.herokuapp.com/api"

	// Initialize with nil, will be set in the init function
	ErrApiData error
	ErrArtists error
)

func init() {
	// Fetch and assign data to ApiData
	ErrApiData = api.FetchData(ApiUrl, &ApiData)
	if ErrApiData == nil {
		ErrArtists = api.FetchData(ApiData["artists"].(string), &MyStruct.Artists)
		if ErrArtists == nil {
			updateArtistData()
		}
	}

	// Run the function after everything is set up

}

// Add the len Key to the aritsts field
func updateArtistData() {
	for i := range MyStruct.Artists {
		if members, ok := MyStruct.Artists[i]["members"].([]interface{}); ok {
			MyStruct.Artists[i]["Len"] = len(members)
		}
	}
}
