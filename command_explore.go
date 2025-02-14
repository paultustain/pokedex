package main

import (
	"fmt"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

func commandExplore(cfg *Config, area_name *string) error {
	fmt.Println(*area_name)

	cachedData, found := cfg.Cache.Get(*area_name)

	if found {
		var cachedPokemon pokeapi.ShallowPokemon
	} else {
		pokemon, err := pokeapi.ListPokemon(*area_name)
	}

	return nil
}
