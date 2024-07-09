package main

import "os"

func commandExit(_ []string, _ playerData) error {
  os.Exit(0)

  return nil
}
