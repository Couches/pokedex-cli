package main

import "fmt"

func commandHelp() error {
  fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
  
  for _, cmd := range getCommands() {
    fmt.Printf("%v: %v \n", cmd.name, cmd.description)
  }

  return nil
}
