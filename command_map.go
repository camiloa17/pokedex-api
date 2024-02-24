package main

import (
	"fmt"
)

func commandMapf(conf *apiConfig) error {
	locations, err := conf.client.ListLocations(conf.next)
	if err != nil {
		return err
	}
	conf.next = locations.Next
	conf.previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(conf *apiConfig) error {
	locations, err := conf.client.ListLocations(conf.previous)
	if err != nil {
		return err
	}
	conf.next = locations.Next
	conf.previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
