package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *apiConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemonInfo, err := conf.client.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	userCatch := rand.Intn(pokemonInfo.BaseExperience)

	if userCatch >= 40 {
		fmt.Printf("%v escaped!\n", name)
		return nil

	}

	conf.pokedex[name] = pokemonInfo
	fmt.Println("you caught " + name)
	return nil
}
