package main

import "fmt"

func commandMapf() error {
  _, locationsResponse := fetchNextLocations()
  locations := locationsResponse.Results
  for _, location := range locations {
    fmt.Println(location.Name)
  }

  return nil
}
