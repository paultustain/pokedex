package pokeapi

type ShallowEncounters struct {
	Location   Location    `json:"location"`
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Location struct {
	Name string
	URL  string
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
