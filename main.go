package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", cleanInput("hello  world!"))
	fmt.Printf("%q\n", cleanInput("hello  \nworld!"))
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}
