package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmlakra/pokedex-cli/internal/pokecache"
)

func Get(url string, cache *pokecache.Cache) ([]byte, error) {
	var data []byte
	if cacheData, present := cache.Get(url); present {
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
		cache.Add(url, data)
	}
	return data, nil
}

func GetPokemon(url string, cache *pokecache.Cache) (PokemonData, error) {
	data, err := Get(url, cache)
	if err != nil {
		return PokemonData{}, err
	}

	var pokemonDataJson PokemonData
	err = json.Unmarshal(data, &pokemonDataJson)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonDataJson, nil
}

func GetLocation(url string, cache *pokecache.Cache) (mapData, error) {
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

func GetLocationDetails(url string, cache *pokecache.Cache) (exploreData, error) {
	data, err := Get(url, cache)
	if err != nil {
		return exploreData{}, err
	}

	var exploreDataJson exploreData
	err = json.Unmarshal(data, &exploreDataJson)
	if err != nil {
		return exploreData{}, err
	}
	return exploreDataJson, nil
}
