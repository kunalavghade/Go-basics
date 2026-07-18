# 1. Simple Start

This topic covers the simple start concepts in Go, specifically how to write a basic "Hello, World!" program.

### Code Example (`main.go`)

```go
package  main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

#### Explanation
1. **`package main`**: Tells the Go compiler that this file should compile as an executable program rather than a shared library.
2. **`import "fmt"`**: Imports the format package, which contains functions for formatting text, including printing to the console.
3. **`func main()`**: The entry point of the application. The program starts executing from this function.

## Run
```bash
go run main.go
```
