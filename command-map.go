package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	regionMapData, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, regionName := range regionMapData.Results {
		fmt.Println(regionName.Name)
	}
	config.nextLocationsURL = &regionMapData.Next
	if regionMapData.Previous != "" {
		config.prevLocationsURL = &regionMapData.Previous
	}
	return nil
}

func commandMapb(config *Config) error {
	regionMapData, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}
	for _, regionName := range regionMapData.Results {
		fmt.Println(regionName.Name)
	}
	config.nextLocationsURL = &regionMapData.Next
	if regionMapData.Previous != "" {
		config.prevLocationsURL = &regionMapData.Previous
	}
	return nil
}
