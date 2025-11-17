package main

import (
	"time"

	Pokeapi "github.com/skadoodle1201/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapi := Pokeapi.NewClient(60*time.Millisecond, 5*time.Second)
	startRepl(pokeapi)
}
