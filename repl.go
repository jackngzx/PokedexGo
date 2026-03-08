package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "List the current 20 map names",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "List the previous 20 map names",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explore the Pokemon in the location area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch the Pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a caught Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "List all caught Pokemon",
		callback:    commandPokedex,
	},
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func repl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		word := words[0]
		args := words[1:]
		if cmd, exists := commands[word]; exists {
			err := cmd.callback(c, args...)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
