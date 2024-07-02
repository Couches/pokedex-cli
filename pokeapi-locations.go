package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Couches/pokecache"
)

var mapCache pokecache.Cache
var nextLocationsURL string = "https://pokeapi.co/api/v2/location-area/"
var prevLocationsURL string

func fetchLocations(url string) (error, LocationsResponse) {
  req, err := http.NewRequest("GET", url, nil)

  locationsResponse := LocationsResponse{}
  
  if err != nil {
    return err, locationsResponse
  }

  res, err := http.DefaultClient.Do(req)

  if err != nil {
    return err, locationsResponse
  }

  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)  

  if err != nil {
    return err, locationsResponse
  }

  err = json.Unmarshal(body, &locationsResponse)

  if err != nil {
    return err, locationsResponse
  }

  nextLocationsURL = locationsResponse.Next
  prevLocationsURL = locationsResponse.Previous

  return nil, locationsResponse
}

func fetchNextLocations() (error, LocationsResponse) {
  return fetchLocations(nextLocationsURL)
}

func fetchPreviousLocations() (error, LocationsResponse) {
  return fetchLocations(prevLocationsURL)
}
