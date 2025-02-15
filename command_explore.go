package main

import (
	"encoding/json"
	"fmt"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

func commandExplore(cfg *Config, name *string) error {

	cachedData, found := cfg.Cache.Get(*name)

	if found {
		var cachedPokemon pokeapi.ShallowEncounters

		err := json.Unmarshal(cachedData, &cachedPokemon)
		if err != nil {
			return err
		}

		for _, result := range cachedPokemon.Encounters {
			fmt.Println(result.Pokemon.Name)
		}

	} else {
		pokemon, err := pokeapi.ListExplore(*name)
		if err != nil {
			return err
		}
		data, err := json.Marshal(pokemon)
		if err != nil {
			return err
		}

		cfg.Cache.Add(*name, data)
		for _, result := range pokemon.Encounters {
			fmt.Println(result.Pokemon.Name)
		}
	}

	return nil
}
