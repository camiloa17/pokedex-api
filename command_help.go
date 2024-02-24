package main

import "fmt"

func commandHelp(conf *apiConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, command := range getCommands() {
		help := fmt.Sprintf("%v: %v", command.name, command.description)
		fmt.Println(help)
	}

	return nil

}
