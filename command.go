package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

type config struct {
	Next     *string
	Previous *string
}

type mapData struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Next != nil {
		url = *c.Next
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	mD := mapData{}
	if err := json.Unmarshal(data, &mD); err != nil {
		log.Fatal(err)
	}

	for _, area := range mD.Results {
		fmt.Println(area.Name)
	}

	c.Next = mD.Next
	c.Previous = mD.Previous

	return nil
}

func commandMapb(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Previous != nil {
		url = *c.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	mD := mapData{}
	if err := json.Unmarshal(data, &mD); err != nil {
		log.Fatal(err)
	}

	for _, area := range mD.Results {
		fmt.Println(area.Name)
	}

	c.Next = mD.Next
	c.Previous = mD.Previous

	return nil
}
