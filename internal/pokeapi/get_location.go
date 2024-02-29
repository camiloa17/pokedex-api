package pokeapi

import "encoding/json"

func (c *Client) GetLocation(locationName string) (ExploreLocationAreaResponse, error) {
	url := c.baseUrl + "/location-area/" + locationName
	val, ok := c.cache.Get(url)

	if ok {
		areaResponse := ExploreLocationAreaResponse{}
		err := json.Unmarshal(val, &areaResponse)
		if err != nil {
			return ExploreLocationAreaResponse{}, err
		}
		return areaResponse, nil
	}

	dat, err := c.getContent(url)

	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}

	areaResponse := ExploreLocationAreaResponse{}
	err = json.Unmarshal(dat, &areaResponse)

	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}
	c.cache.Add(url, dat)
	return areaResponse, nil
}
