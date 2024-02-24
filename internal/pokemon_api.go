package pokemon_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type apiConfig struct {
	url      string
	previous *string
	next     *string
	count    int
}

var apiConf = apiConfig{
	url:      "https://pokeapi.co/api/v2/location?limit=20",
	next:     nil,
	previous: nil,
}

type PokemonLocation struct {
	Name string
}

type locationResponse struct {
	Count    int               `json:"count"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []PokemonLocation `json:"results"`
}

func parseBody(r *http.Response) (locationResponse, error) {
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Println(err)
		return locationResponse{}, errors.New("issue reading the body")
	}
	resLocations := locationResponse{}
	err = json.Unmarshal(body, &resLocations)
	if err != nil {
		fmt.Println(err)
		return locationResponse{}, errors.New("issue parsing the body")
	}
	return resLocations, nil
}

func (c *apiConfig) GetNextLocations() ([]PokemonLocation, error) {
	url := c.url
	if c.next != nil {
		url = *c.next
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("issue fetching the request")
	}

	resLocations, err := parseBody(res)
	if err != nil {
		return nil, err
	}

	c.next = resLocations.Next
	c.previous = resLocations.Previous
	return resLocations.Results, nil
}

func (c *apiConfig) GetPrevLocations() ([]PokemonLocation, error) {
	url := c.previous
	if url == nil {
		return nil, errors.New("no previous location")
	}
	res, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("issue fetching the request")
	}

	resLocations, err := parseBody(res)
	if err != nil {
		return nil, err
	}

	c.next = resLocations.Next
	c.previous = resLocations.Previous
	return resLocations.Results, nil
}
