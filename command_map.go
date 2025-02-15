package main

import (
	"encoding/json"
	"fmt"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

func commandMap(cfg *Config, area_name *string) error {
	if cfg.NextLocationURL == "" {
		fmt.Println("you’re on the last page")
		return nil
	}
	cachedData, found := cfg.Cache.Get(cfg.NextLocationURL)

	if found {
		var cachedLocations pokeapi.ShallowLocations
		err := json.Unmarshal(cachedData, &cachedLocations)
		if err != nil {
			return err
		}

		cfg.NextLocationURL = cachedLocations.Next
		cfg.PreviousLocationURL = cachedLocations.Previous

		for _, result := range cachedLocations.Results {
			fmt.Println(result.Name)
		}
		return nil
	} else {
		locations, err := cfg.Client.ListLocations(cfg.NextLocationURL)
		if err != nil {
			return err
		}

		data, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		cfg.Cache.Add(cfg.NextLocationURL, data)

		cfg.NextLocationURL = locations.Next
		cfg.PreviousLocationURL = locations.Previous

		for _, result := range locations.Results {
			fmt.Println(result.Name)
		}

		return nil
	}

}

func commandMapb(cfg *Config, area_name *string) error {
	if cfg.PreviousLocationURL == "" {
		fmt.Println("you’re on the first page")
		return nil
	}

	locations, err := cfg.Client.ListLocations(cfg.PreviousLocationURL)
	if err != nil {
		return err
	}

	cfg.NextLocationURL = locations.Next
	cfg.PreviousLocationURL = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}
