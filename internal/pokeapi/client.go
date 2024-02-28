package pokeapi

import (
	"net/http"
	"time"

	"github.com/camiloa17/pokedex-api/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	baseUrl    string
	cache      pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseUrl: "https://pokeapi.co/api/v2",
		cache:   *pokecache.NewCache(time.Duration(5 * time.Second)),
	}
}
