package main

import (
	"fmt"

	"github.com/nmlakra/pokedex-cli/internal/pokeapi"
)

func commandMap(config *Config) error {
	offset := 0
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%d&offset=%d", config.limit, offset)
	if config.next != "" {
		url = config.next
	}
	mapDataJson, err := pokeapi.GetLocationAreaData(url, &config.cache)
	if err != nil {
		return err
	}

	config.next = mapDataJson.Next
	config.previous = mapDataJson.Previous

	mapResults := mapDataJson.Results
	for _, res := range mapResults {
		fmt.Println(res.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	if config.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.previous
	mapDataJson, err := pokeapi.GetLocationAreaData(url, &config.cache)
	if err != nil {
		return err
	}

	config.next = mapDataJson.Next
	config.previous = mapDataJson.Previous

	mapResults := mapDataJson.Results
	for _, res := range mapResults {
		fmt.Println(res.Name)
	}

	return nil
}
