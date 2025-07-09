package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type mapData struct {
	Count    int    `json:"Count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaData(url string) (mapData, error) {
	res, err := http.Get(url)
	if err != nil {
		return mapData{}, err
	}
	defer res.Body.Close()

	var data mapData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return mapData{}, err
	}

	return data, nil
}

func commandMap(config *Config) error {
	offset := 0
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%d&offset=%d", config.limit, offset)
	if config.next != "" {
		url = config.next
	}
	mapJsonData, err := getLocationAreaData(url)
	if err != nil {
		return err
	}

	config.next = mapJsonData.Next
	config.previous = mapJsonData.Previous

	mapResults := mapJsonData.Results
	for _, res := range mapResults {
		fmt.Println(res.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	if config.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := config.previous
	mapJsonData, err := getLocationAreaData(url)
	if err != nil {
		return err
	}

	config.next = mapJsonData.Next
	config.previous = mapJsonData.Previous

	mapResults := mapJsonData.Results
	for _, res := range mapResults {
		fmt.Println(res.Name)
	}

	return nil
}
