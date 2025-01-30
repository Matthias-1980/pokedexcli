package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
)

func commandExit(conf *Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

