package main

import "fmt"

func commandInspect(args []string, playerData playerData) error {
  if len(args) == 0 {
    fmt.Println("Missing pokemon to inspect!\nUsage: inspect <Pokemon>")
    return nil
  }

  name := args[0]
  pokemon, ok := playerData.pokedex[name]
  if !ok {
    fmt.Println("That Pokemon isn't in your Pokedex!")
    return nil
  }

  fmt.Printf("Name: %v\n", name)
  fmt.Printf("Height: %v\n", pokemon.Height)
  fmt.Printf("Weight: %v\n", pokemon.Weight)
  fmt.Printf("Stats:\n")

  for _, stat := range pokemon.Stats {
    fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
  }

  fmt.Printf("Types:\n")

  for _, t := range pokemon.Types {
    fmt.Printf("  - %v\n", t.Type.Name)
  }

  return nil
}
