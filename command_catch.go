package main

import (
	"fmt"
	"math/rand"

	"github.com/nmlakra/pokedex-cli/internal/pokeapi"
)

func parsePokemonStats(data *pokeapi.PokemonData) PokemonStats {
	var res PokemonStats
	for _, stat := range data.Stats {
		if stat.Stat.Name == "hp" {
			res.HP = stat.BaseStat
		}
		if stat.Stat.Name == "attack" {
			res.Attack = stat.BaseStat
		}
		if stat.Stat.Name == "defense" {
			res.Defense = stat.BaseStat
		}
		if stat.Stat.Name == "special-attack" {
			res.SpecialAttack = stat.BaseStat
		}
		if stat.Stat.Name == "special-defense" {
			res.SpecialDefense = stat.BaseStat
		}
		if stat.Stat.Name == "speed" {
			res.Speed = stat.BaseStat
		}
	}
	return res
}

func parsePokemonTypes(data *pokeapi.PokemonData) []string {
	var typeRes []string
	for _, t := range data.Types {
		typeRes = append(typeRes, t.Type.Name)
	}
	return typeRes
}

func parsePokemonData(data *pokeapi.PokemonData) Pokemon {
	return Pokemon{
		Name:   data.Name,
		Height: data.Height,
		Weight: data.Weight,
		Stats:  parsePokemonStats(data),
		Types:  parsePokemonTypes(data),
	}
}

func commandCatch(config *Config) error {

	pokemonId := config.commandArgs[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonId)

	res, err := pokeapi.GetPokemon(url, &config.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonId)

	var catchchance int
	if res.BaseExperience > 300 {
		catchchance = 10
	} else if res.BaseExperience >= 100 {
		catchchance = 25
	} else {
		catchchance = 50
	}

	if rand.Intn(100) > catchchance {
		fmt.Printf("%s escaped!\n", pokemonId)
	} else {
		fmt.Printf("%s was caught\n", pokemonId)
		config.pokedex[pokemonId] = parsePokemonData(&res)
	}

	return nil
}
