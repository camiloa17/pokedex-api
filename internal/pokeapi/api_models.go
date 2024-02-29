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

type AreaPokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Encounter struct {
	Pokemon AreaPokemon
}
type ExploreLocationAreaResponse struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

type Type struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonStats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Pokemon struct {
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Weight         int            `json:"weight"`
	Height         int            `json:"height"`
	Stats          []PokemonStats `json:"stats"`
	Types          []Types        `json:"types"`
}
