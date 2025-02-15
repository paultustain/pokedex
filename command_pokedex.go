package main

import "fmt"

func commandPokedex(cfg *Config, name *string) error {
	fmt.Println("Your Pokedex: ")
	for key := range cfg.Pokedex {
		fmt.Println(key)
	}

	return nil
}
