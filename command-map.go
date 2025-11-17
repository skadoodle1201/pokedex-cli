package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Pokecache "github.com/skadoodle1201/pokedexcli/internal/pokecache"
)

type RegionMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func getRegionMap(cache *Pokecache.Cache, url string) (RegionMap, error) {
	var bodyBytes []byte

	bodyBytes, ok := cache.Get(url)

	if !ok {
		response, err := http.Get(url)
		if err != nil {
			return RegionMap{}, err
		}
		defer response.Body.Close()
		bodyBytes, err = io.ReadAll(response.Body)
		if err != nil {
			return RegionMap{}, err
		}
		cache.Add(url, bodyBytes)
	}

	regionMapData := RegionMap{}
	err := json.Unmarshal(bodyBytes, &regionMapData)
	if err != nil {
		return RegionMap{}, err
	}

	return regionMapData, nil
}

func commandMap(config *Config) error {

	regionMapData, err := getRegionMap(config.Cache, config.Next)
	if err != nil {
		return err
	}

	for _, regionName := range regionMapData.Results {
		fmt.Println(regionName.Name)
	}
	config.Next = regionMapData.Next
	if regionMapData.Previous != "" {
		config.Previous = regionMapData.Previous
	}
	return nil
}

func commandMapb(config *Config) error {
	regionMapData, err := getRegionMap(config.Cache, config.Previous)
	if err != nil {
		return err
	}
	for _, regionName := range regionMapData.Results {
		fmt.Println(regionName.Name)
	}
	config.Next = regionMapData.Next
	if regionMapData.Previous != "" {
		config.Previous = regionMapData.Previous
	}
	return nil
}
