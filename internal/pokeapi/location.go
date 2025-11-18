package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RegionMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (client *Client) ListLocations(PageURL *string) (RegionMap, error) {
	url := baseURL + "/location-area"

	if PageURL != nil {
		url = *PageURL
	}
	var responseBody []byte

	if cachedData, ok := client.Cache.Get(url); ok {
		fmt.Println("Cache Hit")
		responseBody = cachedData
	} else {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return RegionMap{}, err
		}

		resp, err := client.Client.Do(req)
		if err != nil {
			fmt.Println("Error after client DO", err)
			return RegionMap{}, err
		}

		defer resp.Body.Close()
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error after readall", err)
			return RegionMap{}, err
		}
	}

	var locationData RegionMap

	if err := json.Unmarshal(responseBody, &locationData); err != nil {
		fmt.Println("Error after Unmarshal", err)
		return RegionMap{}, err
	}
	client.Cache.Add(url, responseBody)

	return locationData, nil

}
