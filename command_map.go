package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMap(conf *Config) error {
	fmt.Println("Printing maps")
	//res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if conf.next == "" {
		fmt.Println("You're on the last page")
		return nil
	}
	res, err := http.Get(conf.next) 
	// what if conf.next == "" or Null

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
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

