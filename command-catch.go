package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon requried")
	}
	pokemonName := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %s... \n", pokemonName)
	pokemonData, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("no pokemon found")
	}
	baseScore := pokemonData.BaseExperience
	randomVal := rand.Intn(200)

	if baseScore > randomVal {
		fmt.Println(pokemonName, "escaped!")
	} else {
		config.pokeapiClient.Pokedex[pokemonName] = pokemonData
		fmt.Println(pokemonName, "was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
	}

	return nil
}
