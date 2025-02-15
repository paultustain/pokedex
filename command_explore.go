package main

import (
	"encoding/json"
	"fmt"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

func commandExplore(cfg *Config, area_name *string) error {
	fmt.Println(*area_name)

	cachedData, found := cfg.Cache.Get(*area_name)

	if found {
		var cachedPokemon pokeapi.ShallowPokemon

		err := json.Unmarshal(cachedData, &cachedPokemon)
		if err != nil {
			return err
		}

		for _, result := range cachedPokemon.Encounters {
			fmt.Println(result.Pokemon.Name)
		}

	} else {
		pokemon, err := pokeapi.ListPokemon(*area_name)
		if err != nil {
			return err
		}
		data, err := json.Marshal(pokemon)
		if err != nil {
			return err
		}

		cfg.Cache.Add(*area_name, data)

		for _, result := range pokemon.Encounters {
			fmt.Println(result.Pokemon.Name)
		}
	}

	return nil
}
