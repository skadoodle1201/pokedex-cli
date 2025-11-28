package main

import "fmt"

func commandPokedex(config *Config, args []string) error {
	pokedex := config.pokeapiClient.Pokedex
	fmt.Println("Your Pokedex:")
	for key := range pokedex {
		fmt.Println(" - ", key)
	}
	return nil
}
