package pokeapi

import (
	"net/http"
	"time"

	Pokecache "github.com/skadoodle1201/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache   Pokecache.Cache
	Client  http.Client
	Pokedex map[string]PokemonDetails
}

func NewClient(apiTimeOut time.Duration, cacheTtl time.Duration) Client {
	return Client{
		Cache: Pokecache.NewCache(cacheTtl),
		Client: http.Client{
			Timeout: apiTimeOut,
		},
		Pokedex: map[string]PokemonDetails{},
	}
}
