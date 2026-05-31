package main

import (
	"fmt"
	"os"
	// "os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	if len(args) < 2 {
		fmt.Println("Please provide a file name as an argument.")
		return
	}
	file, _ := os.Open(args[1])
	data := make([]byte, 99999)
	file.Read(data)
	fmt.Println(string(data))
}
