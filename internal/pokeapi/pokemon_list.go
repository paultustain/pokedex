package pokeapi

import "net/http"

func ListPokemon(area string) (ShallowPokemon, error) {
	url := baseURL + "/location-area/" + area

	req, err := http.Get(url)
	if err != nil {
		return ShallowPokemon{}, err
	}
	defer req.Body.Close()

	return nil
}
