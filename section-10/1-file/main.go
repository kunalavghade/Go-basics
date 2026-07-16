package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"io"
)

func main(){
	data := "this is data to write"
	err := os.WriteFile("file_name.txt", []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File written successfully")
	content, err := os.ReadFile("file_name.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File read successfully", string(content))

	// Create FIle 
	file, err := os.Create("file_name1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("this is data to write in file name 1")

	fileread, err := os.Open("file_name1.txt")
	defer fileread.Close()
	scanner := bufio.NewScanner(fileread)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	
}