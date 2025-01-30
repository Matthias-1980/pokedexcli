package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"pokedexcli/internal/pokecache"
)

func commandMap(conf *Config, cache *pokecache.Cache) error {
	fmt.Println("Printing maps")
	//res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if conf.next == "" {
		fmt.Println("You're on the last page")
		return nil
	}

	var body []byte
	var err error

	cacheVal, exists := cache.Get(conf.next)
	if exists {
		fmt.Println("cache exists.")
		body = cacheVal 
	} else {
		fmt.Println("cache does not exist.")
		res, err := http.Get(conf.next) 
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			//return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
        	}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			//return err
		}		

		cache.Add(conf.next, body)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	//setting next pages
	(*conf).next = response.Next
	(*conf).prev = response.Previous

	//fmt.Println(response)
	//fmt.Printf("%s", body)
	for _,names := range response.Results {
		fmt.Println(names.Name)
	}
	return nil
}

