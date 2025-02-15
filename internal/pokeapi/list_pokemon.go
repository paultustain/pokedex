package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func ListPokemon(name string) (ShallowPokemon, error) {
	url := baseURL + "/pokemon/" + name

	req, err := http.Get(url)

	if err != nil {
		return ShallowPokemon{}, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return ShallowPokemon{}, err
	}

	pkmn := ShallowPokemon{}
	err = json.Unmarshal(res, &pkmn)
	if err != nil {
		return ShallowPokemon{}, nil
	}

	// fmt.Println(pkmn)
	return pkmn, nil
}
