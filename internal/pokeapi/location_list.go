package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type mapData struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func (c *Client) ListLocations(pageURL *string) (mapData, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		mD := mapData{}
		err := json.Unmarshal(val, &mD)
		if err != nil {
			return mapData{}, err
		}
		return mD, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return mapData{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return mapData{}, err
	}

	mD := mapData{}
	if err := json.Unmarshal(data, &mD); err != nil {
		return mapData{}, err
	}
	c.cache.Add(url, data)
	return mD, nil
}
