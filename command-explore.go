package main

import "fmt"

func commandExplore(args []string) error {
  location := args[0]
  fmt.Printf("Exploring %s...\n", location)

  err, exploreResponse := fetchLocationPokemon(location)

  if err != nil {
    fmt.Println("I don't think that area exists.")
    return err
  }

  encounters := exploreResponse.PokemonEncounters
  fmt.Println("Found Pokemon:")
  for _, encounter := range encounters {
    fmt.Printf(" - %s\n", encounter.Pokemon.Name)
  }

  return nil
} 
