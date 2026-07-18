package main

import (
	"bytes"
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
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(jane); err != nil {
		log.Println(err)
	}

	fmt.Printf(buf.String())
}