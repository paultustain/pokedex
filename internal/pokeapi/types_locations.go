package pokeapi

type ShallowLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Area `json:"results"`
}

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
