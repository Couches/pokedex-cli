package main

import "fmt"

func commandMapb(_ []string) error {
  _, locationsResponse := fetchPreviousLocations()
  locations := locationsResponse.Results
  for _, location := range locations {
    fmt.Println(location.Name)
  }

  return nil
}
