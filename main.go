package main

import (
	"time"

	"github.com/camiloa17/pokedex-api/internal/pokeapi"
)

func main() {
	config := apiConfig{
		client: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(&config)
}
