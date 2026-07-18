package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
	var payload = `{"name":"jane","age":24,"phone":"34567-890-567","is_active":true}`
	var u user
	buf := strings.NewReader(payload)
	dec := json.NewDecoder(buf)
	if err := dec.Decode(&u); err != nil {
		log.Println(err)
	}

	fmt.Printf("+%v", u)
}