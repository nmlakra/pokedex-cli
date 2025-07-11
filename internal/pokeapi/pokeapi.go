package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nmlakra/pokedex-cli/internal/pokecache"
)

func Get(url string, cache *pokecache.Cache) ([]byte, error) {
	var data []byte
	if cacheData, present := cache.Get(url); present {
		fmt.Println("using cached data!")
		data = cacheData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println("caching data!")
		cache.Add(url, data)
	}
	return data, nil
}

func GetLocationAreaData(url string, cache *pokecache.Cache) (mapData, error) {
	data, err := Get(url, cache)
	if err != nil {
		return mapData{}, err
	}

	var mapDataJson mapData
	err = json.Unmarshal(data, &mapDataJson)
	if err != nil {
		return mapData{}, err
	}
	return mapDataJson, nil
}
