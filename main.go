package main

import (
	"time"

	Pokeapi "github.com/skadoodle1201/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := Pokeapi.NewClient(5*time.Second, 60*time.Millisecond)
	cfg := &Config{
		pokeapiClient: pokeapiClient,
	}
	startRepl(cfg)
}
