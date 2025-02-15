package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL string) (ShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != "" {
		url = pageURL
	}
	req, err := http.Get(url)
	if err != nil {
		return ShallowLocations{}, fmt.Errorf("error creating request: %w", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return ShallowLocations{}, fmt.Errorf("error reading response: %w", err)
	}

	cfg := ShallowLocations{}
	err = json.Unmarshal(res, &cfg)

	if err != nil {
		return ShallowLocations{}, fmt.Errorf("fail to unmarshal: %w", err)
	}

	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}

	return cfg, nil
}
