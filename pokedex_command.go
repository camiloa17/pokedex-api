package main

import "fmt"

func pokedexCommand(conf *apiConfig, args ...string) error {
	for _, pokemon := range conf.pokedex {
		fmt.Printf("- %v \n", pokemon.Name)
	}
	return nil
}
