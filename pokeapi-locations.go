package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
  "fmt"

	"github.com/Couches/pokecache"
)

var mapCache pokecache.Cache = pokecache.NewCache(10 * time.Second)
var locationCache pokecache.Cache = pokecache.NewCache(10 * time.Second)

var nextLocationsURL string = "https://pokeapi.co/api/v2/location-area/"
var prevLocationsURL string

func fetchLocations(url string) (error, LocationsResponse) {
	locationsResponse := LocationsResponse{}

	body, ok := mapCache.Get(url)
  
	if !ok {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return err, locationsResponse
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			return err, locationsResponse
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if err != nil {
			return err, locationsResponse
		}

    mapCache.Add(url, body)
	}

  err := json.Unmarshal(body, &locationsResponse)

	if err != nil {
		return err, locationsResponse
	}

	nextLocationsURL = locationsResponse.Next
	prevLocationsURL = locationsResponse.Previous

	return nil, locationsResponse
}

func fetchLocationPokemon(location string) (error, ExploreResponse) {
  exploreResponse := ExploreResponse{}

  body, ok := locationCache.Get(location)

  if !ok {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
      return err, exploreResponse
    }

    res, err := http.DefaultClient.Do(req)

    if err != nil {
      return err, exploreResponse
    }

    defer res.Body.Close()

    body, err = io.ReadAll(res.Body)

    if err != nil {
      return err, exploreResponse
    }

    locationCache.Add(location, body)
  }

  err := json.Unmarshal(body, &exploreResponse)

  if err != nil {
    return err, exploreResponse
  }

  return nil, exploreResponse
}

func fetchNextLocations() (error, LocationsResponse) {
	return fetchLocations(nextLocationsURL)
}

func fetchPreviousLocations() (error, LocationsResponse) {
	return fetchLocations(prevLocationsURL)
}
