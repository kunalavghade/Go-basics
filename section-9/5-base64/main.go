package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Hello, Gophers!"

	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}
