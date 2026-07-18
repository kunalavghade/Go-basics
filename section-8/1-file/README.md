# 1. File Operations

This topic covers advanced file operations, reading, and writing using the `os` and `bufio` packages.

### Code Example (`main.go`)

```go
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

	// Create File 
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
```

#### Explanation
1. **`os.WriteFile` & `os.ReadFile`**: High-level utility functions to quickly write or read an entire file into memory as a byte slice. `0644` is the standard file permission mode.
2. **`os.Create` & `WriteString`**: Opens a file (creates if it doesn't exist) and provides a `*os.File` pointer, which can be used to write strings sequentially.
3. **`bufio.Scanner`**: A highly efficient way to read a file line-by-line. The `for scanner.Scan()` loop continues until EOF. `scanner.Text()` returns the current line as a string.
4. **`defer file.Close()`**: The `defer` keyword delays the execution of a function until the surrounding function returns. This is the idiomatic way to ensure resources like files are properly closed after opening.

## Run
```bash
go run main.go
```
