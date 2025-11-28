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

type LocationAreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (client *Client) GetLocationInformation(LocationAreaName string) (LocationAreaDetails, error) {
	var locationDetails LocationAreaDetails
	if LocationAreaName == "" {
		return LocationAreaDetails{}, fmt.Errorf("location name required")
	}
	resBody, ok := client.Cache.Get(LocationAreaName)

	if !ok {
		url := baseURL + "/location-area" + "/" + LocationAreaName
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		res, err := client.Client.Do(req)

		if err != nil {
			return LocationAreaDetails{}, err
		}

		defer res.Body.Close()

		resBody, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error after readall", err)
			return LocationAreaDetails{}, err
		}

		if err = json.Unmarshal(resBody, &locationDetails); err != nil {
			return LocationAreaDetails{}, err
		}
	}

	client.Cache.Add(LocationAreaName, resBody)

	return locationDetails, nil
}
