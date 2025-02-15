package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func ListExplore(area string) (ShallowEncounters, error) {
	url := baseURL + "/location-area/" + area + "/"

	req, err := http.Get(url)

	if err != nil {
		return ShallowEncounters{}, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return ShallowEncounters{}, nil
	}

	pkmn := ShallowEncounters{}

	err = json.Unmarshal(res, &pkmn)
	if err != nil {
		return ShallowEncounters{}, nil
	}

	return pkmn, nil
}
