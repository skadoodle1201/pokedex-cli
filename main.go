package main

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Next     string
	Previous string
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
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
