package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

func commandCatch(cfg *Config, name *string) error {
	fmt.Printf("Throwing a Pokeball at %s...", *name)
	pokemon, err := pokeapi.ListPokemon(*name)
	if err != nil {
		return err
	}

	chance := 49 / float64(pokemon.BaseExperience)
	r := rand.Float64()

	if r < chance {
		fmt.Printf("%s was caught!\n", *name)
		cfg.Pokedex[*name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", *name)
	}

	return nil
}
