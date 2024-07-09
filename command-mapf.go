package main

import "fmt"

func commandMapf(_ []string, _ playerData) error {
  _, locationsResponse := fetchNextLocations()
  locations := locationsResponse.Results
  for _, location := range locations {
    fmt.Println(location.Name)
  }

  return nil
}
