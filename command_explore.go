package main

import (
	"fmt"

	"github.com/nmlakra/pokedex-cli/internal/pokeapi"
)

func commandExplore(config *Config) error {
	location := config.commandArgs[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	res, err := pokeapi.GetLocationDetails(url, &config.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
