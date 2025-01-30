package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"pokedexcli/internal/pokecache"
)

func commandMapb(conf *Config, cache *pokecache.Cache) error {
	fmt.Println("Printing maps")
	if conf.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var body []byte
	var err error

	cacheVal, exists := cache.Get(conf.prev)
	if exists {
		fmt.Println("using cache")
		body = cacheVal // "body" declared and not used
	} else {
		fmt.Println("not using cache")
		res, err := http.Get(conf.prev) 
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

		cache.Add(conf.prev, body)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	//setting next and previous pages
	(*conf).next = response.Next
	(*conf).prev = response.Previous

	//fmt.Println(response)
	//fmt.Printf("%s", body)
	for _,names := range response.Results {
		fmt.Println(names.Name)
	}
	return nil
}

