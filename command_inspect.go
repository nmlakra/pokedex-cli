package main

import "fmt"

func commandInspect(config *Config) error {
	pokemonId := config.commandArgs[0]
	pokemon, present := config.pokedex[pokemonId]
	if !present {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats: ")
	fmt.Printf("  -hp: %d\n", pokemon.Stats.HP)
	fmt.Printf("  -attack: %d\n", pokemon.Stats.Attack)
	fmt.Printf("  -defense: %d\n", pokemon.Stats.Defense)
	fmt.Printf("  -special-attack: %d\n", pokemon.Stats.SpecialAttack)
	fmt.Printf("  -special-defense: %d\n", pokemon.Stats.SpecialDefense)
	fmt.Printf("  -speed: %d\n", pokemon.Stats.Speed)
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}
	return nil
}
