package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := Config{}
	config.prev = ""
	config.next = "https://pokeapi.co/api/v2/location-area/"
	conf := &config

	reader := bufio.NewScanner(os.Stdin)
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
			err := command.callback(conf)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map":	{
			name:	"map",
			description: "Shows next 20 maps",
			callback: commandMap,
		},
		"mapb":	{
			name:	"mapb",
			description: "Shows previous 20 maps",
			callback: commandMapb,
		},
	}
}

type Config struct {
	prev	string
	next	string
}

