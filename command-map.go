package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func getRegionMap(url string) (RegionMap, error) {
	response, err := http.Get(url)
	if err != nil {
		return RegionMap{}, err
	}
	defer response.Body.Close()

	var regionMapData RegionMap
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&regionMapData)
	if err != nil {
		return RegionMap{}, err
	}

	return regionMapData, nil
}

func commandMap(config *Config) error {

	regionMapData, err := getRegionMap(config.Next)
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
	regionMapData, err := getRegionMap(config.Previous)
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
