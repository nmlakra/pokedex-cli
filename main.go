package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nmlakra/pokedex-cli/internal/pokecache"
)

type CommandCallback func(config *Config) error

type cliCommand struct {
	name        string
	description string
	callback    CommandCallback
}

type Config struct {
	limit       int
	next        string
	previous    string
	cache       pokecache.Cache
	commandArgs []string
}

func main() {

	var config Config
	config.limit = 20
	config.cache = *pokecache.NewCache(300 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputTexts := cleanInput(scanner.Text())
		commandName := inputTexts[0]
		config.commandArgs = inputTexts[1:]
		if command, ok := validCommands()[commandName]; !ok {
			fmt.Println("Unkown command")
		} else {
			command.callback(&config)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func validCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemon found in the provided Location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
	}
}
