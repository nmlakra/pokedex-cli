package main

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  PokemonStats
	Types  []string
}

type PokemonStats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}
