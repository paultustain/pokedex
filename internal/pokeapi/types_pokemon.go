package pokeapi

type ShallowPokemon struct {
	Results []Pokemon `json:"pokemon_encounters"`
}

type Pokemon struct {
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}
