package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Pokemons []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNumber int    `json:"entry_number"`
	Specie      Specie `json:"pokemon_species"`
}

type Specie struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main1() {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(res.Body) // byteArr
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	var response Response

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Fatal(err)
	}

	for i, pokemon := range response.Pokemons {
		fmt.Printf("Pokemon %d is %s\n", i+1, pokemon.Specie.Name)
	}
}
