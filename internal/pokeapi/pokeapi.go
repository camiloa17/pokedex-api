package pokeapi

import (
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
