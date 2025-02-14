package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/paultustain/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *string) error
}

type Config struct {
	Client              pokeapi.Client
	Cache               pokeapi.Cache
	NextLocationURL     string `json:"next"`
	PreviousLocationURL string `json:"previous"`
}

func main() {
	interval := 5 * time.Second

	config := Config{
		Cache:           pokeapi.NewCache(interval),
		NextLocationURL: "https://pokeapi.co/api/v2/location-area/",
	}

	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help with the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Map the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Map the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon in an area",
			callback:    commandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		command := cleanInput(scanner.Text())[0]
		var area_name string
		if command == "explore" {
			if len(cleanInput(scanner.Text())) < 2 {
				fmt.Printf("Please add area name to explore.\n")
				continue
			} else {
				area_name = cleanInput(scanner.Text())[1]
			}

		}

		item, ok := cliCommands[command]
		if ok {
			item.callback(&config, &area_name)
		} else {
			fmt.Printf("Unkown Command\n")
		}
	}
}

func cleanInput(text string) []string {
	lowerS := strings.ToLower(text)
	splitS := strings.Fields(lowerS)
	return splitS

}
