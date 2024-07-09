package main

import "fmt"
import "math/rand/v2"
import "time"

func commandCatch(args []string, playerData playerData) error {
  if len(args) == 0 {
    fmt.Println("Missing pokemon to catch!\nUsage: catch <Pokemon>")
    return nil
  }
  target := args[0]

  err, pokemon := fetchPokemon(target)

  if err != nil {
    fmt.Println("I don't think that Pokemon exists...")
    return nil
  }

  name := pokemon.Name

  fmt.Printf("Throwing a Pokeball at %v...\n", name)

  time.Sleep(1 * time.Second)

  chance := pokemon.BaseExperience / 30
  roll := randRange(1, chance)

  if roll == chance {
    fmt.Printf("%s was caught!\n", name)
    playerData.pokedex[name] = pokemon
  } else {
    fmt.Printf("%s escaped!\n", name)
  }

  return nil
}


func randRange(min, max int) int {
  return rand.IntN(max + 1 - min) + min
}
