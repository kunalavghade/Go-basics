package main

import (
	"embed"
	"fmt"
	// "io/fs"
	"log"
)

//go:embed *.txt
var data string


//go:embed public
var publc embed.FS

func main() {
	fmt.Println(data)

	// data, err := fs.ReadFile(publc, "data.txt")
	data, err := publc.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	

}