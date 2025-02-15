package pokeapi

type ShallowPokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stat  `json:"stats"`
	Types          []Type  `json:"types"`
	BaseExperience float64 `json:"base_experience"`
}

type Stat struct {
	Information StatInformation `json:"stat"`
	BaseStat    int             `json:"base_stat"`
}

type StatInformation struct {
	Name string `json:"name"`
}

type Type struct {
	Information TypeInformation `json:"type"`
}

type TypeInformation struct {
	Name string `json:"name"`
}
