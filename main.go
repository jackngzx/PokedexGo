package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := strings.Fields(strings.ToLower(text))
		word := words[0]
		if cmd, exists := commands[word]; exists {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
