package main

import "fmt"

func commandPokedex(config *Config) error {
	if len(config.pokedex) == 0 {
		fmt.Println("no pokemon caught!")
	}
	fmt.Println("Your Pokedex: ")
	for key := range config.pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
