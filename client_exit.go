package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, area_name *string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
