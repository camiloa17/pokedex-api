package main

import (
	"errors"
	"fmt"
)

func inspectCommand(conf *apiConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]

	pokemon, ok := conf.pokedex[name]

	if !ok {
		fmt.Printf("You have not caught %v yet \n", name)
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ts := range pokemon.Types {
		fmt.Printf(" -%v\n", ts.Type.Name)
	}

	return nil
}
