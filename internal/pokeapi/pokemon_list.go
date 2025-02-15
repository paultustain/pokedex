package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func ListPokemon(area string) (ShallowPokemon, error) {
	url := baseURL + "/location-area/" + area + "/"

	req, err := http.Get(url)

	if err != nil {
		return ShallowPokemon{}, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return ShallowPokemon{}, nil
	}

	pkmn := ShallowPokemon{}

	err = json.Unmarshal(res, &pkmn)
	if err != nil {
		return ShallowPokemon{}, nil
	}

	return pkmn, nil
}
