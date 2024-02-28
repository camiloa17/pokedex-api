package main

import (
	"errors"
	"fmt"
)

func explore(conf *apiConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	areaInfo, err := conf.client.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Println(areaInfo)
	for _, encounter := range areaInfo.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
