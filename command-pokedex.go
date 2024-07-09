package main

import "fmt"

func commandPokedex(_ []string, playerData playerData) error {
  fmt.Println("Your Pokedex:")

  for _, pokemon := range playerData.pokedex {
    fmt.Printf("  - %v\n", pokemon.Name)
  }

  return nil
}
