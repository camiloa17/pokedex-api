package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationResponse, error) {
	url := c.baseUrl + "/location?limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		return LocationResponse{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationResponse{}, err
	}

	locationsResp := LocationResponse{}
	err = json.Unmarshal(dat, &locationsResp)

	if err != nil {
		return LocationResponse{}, err
	}

	return locationsResp, nil
}
