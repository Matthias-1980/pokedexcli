package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMapb(conf *Config) error {
	fmt.Println("Printing maps")
	if conf.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(conf.prev) 
	// what if conf.prev == "" or Null

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

