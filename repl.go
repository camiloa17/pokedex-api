package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/camiloa17/pokedex-api/internal/pokeapi"
)

type apiConfig struct {
	previous *string
	next     *string
	client   pokeapi.Client
}

func startRepl(config *apiConfig) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		word := cleanInput(reader.Text())
		if len(word) == 0 {
			continue
		}

		commandName := word

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) string {
	output := strings.ToLower(text)
	words := strings.TrimSpace(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*apiConfig) error
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
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}
