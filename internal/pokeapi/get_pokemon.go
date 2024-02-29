package pokeapi

import "encoding/json"

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := c.baseUrl + "/pokemon/" + pokemonName
	val, ok := c.cache.Get(url)
	if ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	dat, err := c.getContent(url)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResponse := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResponse)

	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)
	return pokemonResponse, nil
}
