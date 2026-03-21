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
	callback    func() error
}

type Response struct {
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

var data Response

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		cmd, exits := getCommands()[commandName]
		if exits {
			err := cmd.callback()
			if err != nil {
				fmt.Println("The error :", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display all locations",
			callback: func() error {
				var url string
				if data.Next == "" {
					url = "https://pokeapi.co/api/v2/location-area/"
				} else {
					url = data.Next
				}
				return commandMap(url)
			},
		},
		"mapb": {
			name:        "map",
			description: "Display all locations",
			callback: func() error {
				var url string
				if data.Previous == "" {
					url = "https://pokeapi.co/api/v2/location-area/"
				} else {
					url = data.Previous
				}
				return commandMap(url)
			},
		},
	}
}
