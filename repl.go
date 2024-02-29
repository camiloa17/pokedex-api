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
	pokedex  map[string]pokeapi.Pokemon
}

func startRepl(config *apiConfig) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(apiConfig *apiConfig, args ...string) error
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
		"explore": {
			name:        "explore",
			description: "Get the information of pokemon in a specific area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects the pokedex for your pokemons information",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the pokemons in your pokedex",
			callback:    pokedexCommand,
		},
	}
}
