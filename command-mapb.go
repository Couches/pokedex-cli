package main

import "fmt"

func commandMapb() error {
  _, locationsResponse := fetchPreviousLocations()
  locations := locationsResponse.Results
  for _, location := range locations {
    fmt.Println(location.Name)
  }

  return nil
}
