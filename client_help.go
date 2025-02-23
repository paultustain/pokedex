package main

import "fmt"

func commandHelp(cfg *Config, name *string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
catch: Attempts to catch a pokemon
map: Returns a list of 20 locations at a time
mapb: Goes back to the previous 20 lcoations
Explore: Lists the available Pokemon in the area`)
	return nil
}
