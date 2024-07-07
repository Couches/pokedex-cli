package main

import "fmt"

func commandCatch(args []string) error {
  target := args[0]

  err, pokemon := fetchPokemon(target)

  if err != nil {
    fmt.Println("I don't think that Pokemon exists...")
  }

  fmt.Println(pokemon)

  return nil
}
