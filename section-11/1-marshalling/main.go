package main

import (
	"encoding/json"
	"fmt"
	"log"
)
// user struct represents the data we want to marshal into JSON.
// Struct tags like `json:"name"` dictate the key names in the resulting JSON.
type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func main() {
	jane := user {
		Name: "jane",
		Age: 23,
		IsActive: true,
		Phone: "34567-890-567",
	}
	// json.Marshal converts a Go data structure (like our user struct) into a JSON-encoded byte slice.
	// The output is compact, without spaces or newlines.
	byteSlice, err := json.Marshal(jane)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(byteSlice))

	// json.MarshalIndent works similarly but formats the JSON for readability.
	// The second argument is a prefix for each line (here "-"), and the third is the indent (here " ").
	byteFormated, err := json.MarshalIndent(jane, "-", " ")
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(byteFormated))
}