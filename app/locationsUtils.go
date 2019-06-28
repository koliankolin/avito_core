package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)


type Location struct {
	Локации map[string][]string
}

type Locations struct {
	Россия Location
}


func getAllLocations(letter string) []byte {
	resp, _ := http.Get("https://www.avito.ru/web/1/slocations?limit=10000&q=" + letter)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func getSubLocations() (*Locations, error) {
	sublocations := new(Location)
	sublocations.Локации = make(map[string][]string)
	finSubLocations := new(Locations)

	var result map[string]map[string][]interface{}

	for letter := rune('а'); string(letter) != "я"; letter++ {
		locations := getAllLocations(string(letter))
		err := json.Unmarshal(locations, &result)
		if err != nil {
			return nil, nil
		}

		for _, val := range result["result"]["locations"] {
			name := val.(map[string]interface{})["names"].(map[string]interface{})["1"].(string)
			parent := val.(map[string]interface{})["parent"]
			if parent != nil {
				parentName := parent.(map[string]interface{})["names"].(map[string]interface{})["1"].(string)
				parentName = parentName
				sublocations.Локации[parentName] = append(sublocations.Локации[parentName], name)
			}
		}
	}

	// Add Moscow and Piter
	sublocations.Локации["Москва"] = append(sublocations.Локации["Москва"], "Москва")
	sublocations.Локации["Санкт-Петербург"] = append(sublocations.Локации["Санкт-Петербург"], "Санкт-Петербург")

	finSubLocations.Россия = *sublocations

	return finSubLocations, nil
}

func getSublocationsJson() string {
	subLocations, err := getSubLocations()

	if err != nil {
		log.Fatal(err)
	}

	subLocationsJson, err := json.Marshal(subLocations)

	if err != nil {
		log.Fatal(err)
	}

	return string(subLocationsJson)
}

func getSublocationsOne(location string) string {
	arr := strings.Split(location, " ")
	if len(arr) > 2 {
		arr[0] = strings.Title(arr[0])
	}
	location = strings.Join(arr, " ")

	locations, err := getSubLocations()

	if err != nil {
		log.Fatal(err)
	}
	if cities, ok := locations.Россия.Локации[location]; ok == true {

		loc := new(Location)
		loc.Локации = make(map[string][]string)
		loc.Локации[location] = cities

		locJson, err := json.Marshal(loc)
		if err != nil {
			log.Fatal(err)
		}

		return string(locJson)
	}

	return `{"message":"No such location"}`
}

func printHumanReadableOne(location string) error {
	arr := strings.Split(location, " ")
	arr[0] = strings.Title(arr[0])
	location = strings.Join(arr, " ")

	locations, err := getSubLocations()

	if err != nil {
		return err
	}
	if cities, ok := locations.Россия.Локации[location]; ok == true {
		fmt.Println(location)
		if len(cities) - 1 >= 0 {
			for _, city := range cities[:len(cities) - 1] {
				fmt.Printf("\t├── %s\n", city)
			}
		}

		fmt.Printf("\t└── %s\n", cities[len(cities) - 1])
		return nil
	} else {
		return nil
	}
}

func printHumanReadableAll() error {
	locations, err := getSubLocations()
	keys := reflect.ValueOf(locations.Россия.Локации).MapKeys()

	fmt.Println(locations)

	if err != nil {
		return err
	}

	fmt.Println("Россия")

	for _, key := range keys[:len(keys) - 1]{
		cities := locations.Россия.Локации[key.String()]
		fmt.Printf("\t├── %s\n", key.String())
		if len(cities) - 1 >= 0 {
			for _, city := range cities[:len(cities) - 1] {
				fmt.Printf("\t\t├── %s\n", city)
			}
		}

		fmt.Printf("\t\t└── %s\n", cities[len(cities) - 1])
	}
	return nil
}





