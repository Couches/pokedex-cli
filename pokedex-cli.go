package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("There was an error reading your input.")
			return
		}

    input = cleanInput(input)

		if cmd, exists := getCommands()[input]; exists {
			cmd.callback()
      fmt.Print("\n")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
      callback: commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
      callback: commandExit,
		},
    "map": {
      name: "map",
      description: "Lists the next 20 location areas",
      callback: commandMapf,
    },
    "mapb": {
      name: "mapb",
      description: "Lists the previous 20 location areas",
      callback: commandMapb,
    },
	}
}

func cleanInput(input string) string {
  input = strings.TrimSpace(input)
  input = strings.ToLower(input)

  return input
}
