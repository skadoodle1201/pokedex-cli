package main

import "fmt"

func commandExplore(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid command area not specified")
	}

	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")

	areaName := args[0]

	locationDetails, err := config.pokeapiClient.GetLocationInformation(areaName)

	if err != nil {
		return err
	}

	pokemonsFound := locationDetails.PokemonEncounters

	for _, val := range pokemonsFound {
		fmt.Println(" - ", val.Pokemon.Name)
	}

	return nil
}
