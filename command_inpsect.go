package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, name *string) error {

	information, found := cfg.Pokedex[*name]

	if !found {
		fmt.Printf("Not caught %s yet. \n", *name)
		return errors.New("not caught this yet")
	}

	fmt.Printf("Name: %s\n", information.Name)
	fmt.Printf("Height: %d\n", information.Height)
	fmt.Printf("Weight: %d\n", information.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range information.Stats {
		fmt.Printf("- %s: %d\n", stat.Information.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, pokeType := range information.Types {
		fmt.Printf("- %s\n", pokeType.Information.Name)
	}
	return nil
}
