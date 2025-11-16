package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type regionMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func commandMap() error {
	response, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var regionMapData regionMap
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&regionMapData)
	if err != nil {
		return err
	}

	for _, regionName := range regionMapData.Results {
		fmt.Println(regionName.Name)
	}

	return nil

}
