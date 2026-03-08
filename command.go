package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/jackngzx/PokedexGo/internal/pokeapi"
)

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: List the next 20 location areas")
	fmt.Println("mapb: List the previous 20 location areas")
	fmt.Println("explore: Explore and find Pokemon inside the provided location area")
	fmt.Println("catch: Catch a Pokemon")
	fmt.Println("inspect: Inspect stats of a caught Pokemon")
	fmt.Println("pokedex: List all caught Pokemon")
	return nil
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokedex       map[string]pokeapi.PokemonData
}

func commandMap(c *config, args ...string) error {
	location, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}
	c.Next = location.Next
	c.Previous = location.Previous

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(c *config, args ...string) error {
	location, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return err
	}
	c.Next = location.Next
	c.Previous = location.Previous

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(c *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide location name")
	}
	locationName := args[0]
	locationArea, err := c.pokeapiClient.LocationExplore(locationName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil

}

func commandCatch(c *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide Pokemon name")
	}
	pokemon, err := c.pokeapiClient.PokemonDataGet(args[0])
	if err != nil {
		return fmt.Errorf("Pokemon could not be found in the database. Try again!")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchChance := rand.Intn(pokemon.BaseExperience)
	if catchChance > 40 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command")
		c.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(c *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide caught Pokemon name")
	}
	name := args[0]
	if _, ok := c.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", c.pokedex[name].Name)
		fmt.Printf("Height: %d\n", c.pokedex[name].Height)
		fmt.Printf("Weight: %d\n", c.pokedex[name].Weight)
		fmt.Println("Stats:")
		for _, s := range c.pokedex[name].Stats {
			fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range c.pokedex[name].Types {
			fmt.Printf(" -%s\n", t.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("Pokemon not yet caught")
}

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		return fmt.Errorf("You have not caught any Pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for _, p := range c.pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
