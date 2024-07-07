package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("There was an error reading your input.")
			return
		}

		args := strings.Fields(cleanInput(input))

		if cmd, exists := getCommands()[args[0]]; exists {
			cmd.callback(args[1:])
			fmt.Print("\n")
		} else {
      fmt.Println("I don't think that command exists.")
      commandHelp(args[1:])
    }
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "Lists all pokemon that may be found in an area",
			callback:    commandExplore,
		},
    "catch": {
      name: "catch <Pokemon>",
      description: "Throws a ball at the Pokemon",
      callback: commandCatch,
    },
	}
}

func cleanInput(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input
}
