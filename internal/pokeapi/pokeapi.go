package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return dat, nil
}

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
