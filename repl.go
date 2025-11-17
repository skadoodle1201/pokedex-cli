package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cache := Pokecache.NewCache(15 * time.Second)

	config := Config{

		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
		Cache:    cache,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}
