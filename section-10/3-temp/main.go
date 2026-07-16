package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	tmpfile , err := os.CreateTemp("", "tmpfile.txt")
	if err != nil{
		log.Fatal("error creating temp file: ", err)
	}
	defer func() {
		fmt.Println("Removeing File", tmpfile.Name())
		os.Remove(tmpfile.Name())
	}()

	_, err = tmpfile.Write([]byte("hello from temp file"))
	if err != nil{
		log.Fatal("Error Writing to temp File: ", err)
	}
	fmt.Println("Write to temp file success : ", tmpfile.Name())

	tmpDir, err := os.MkdirTemp("", "tmpdir*")
	if err != nil{
		log.Fatal("error creating temp dir: ", err)
	}
	fmt.Println("Created temp dir: ", tmpDir)
	defer func() {
		fmt.Println("Removeing Dir", tmpDir)
		os.RemoveAll(tmpDir)
	}()

}