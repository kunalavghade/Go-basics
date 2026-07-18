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
	Role	 string `json:"role"`
	Profile  profile `json:"profile"`
}

type profile struct {
	Address string `json:"address"`
	Website string `json:"website"`
}

var payload = `{
	"name": "john",
	"age": 30,
	"phone": "1234567890",
	"is_active": true,
	"profile":{
		"address": "123 Main St",
		"website": "www.example.com"
	}
}`
func main() {
	var u user
	err := json.Unmarshal([]byte(payload), &u)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v",u)
}