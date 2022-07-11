package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "sai",
		"last_name": "kumar",
		"hair_color": "black",
		"has_dog": false
	},
	{
		"first_name": "hemmenth",
		"last_name": "naidu",
		"hair_color": "black",
		"has_dog": true
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error marshalling json", err)
	}
	log.Printf("unmarshalle: %v", unmarshalled)

	//write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "lokesh"
	m1.LastName = "damma"
	m1.HairColor = "black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "sri"
	m2.LastName = "kanth"
	m2.HairColor = "black"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))

}
