package pokeapi

type PokemonLocationArea struct {
	Name string `json:"name"`
}

type LocationAreasResponse struct {
	Count    int                   `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []PokemonLocationArea `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Encounter struct {
	Pokemon Pokemon
}
type ExploreLocationAreaResponse struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}
