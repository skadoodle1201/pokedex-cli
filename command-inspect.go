package main

import (
	"fmt"
	"strings"
)

func commandInspect(config *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid args")
	}

	pokemonName := strings.ToLower(args[0])

	pokemon, ok := config.pokeapiClient.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s \nHeight: %v \nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for key := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", pokemon.Stats[key].Stat.Name, pokemon.Stats[key].BaseStat)
	}
	fmt.Println("Types:")
	for key := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon.Types[key].Type.Name)
	}

	return nil
}
