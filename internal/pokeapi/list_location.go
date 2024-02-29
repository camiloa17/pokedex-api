package pokeapi

import "encoding/json"

func (c *Client) ListLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := c.baseUrl + "/location-area?limit=20"

	if pageUrl != nil {
		url = *pageUrl
	}
	val, ok := c.cache.Get(url)

	if ok {
		locationResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationResp, nil
	}

	dat, err := c.getContent(url)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationsResp := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationsResp)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
