# 3. FileReader

This topic covers reading files in Go from command line arguments using interfaces.

### Code Example (`main.go`)

```go
package main

import (
	"fmt"
	"os"
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
```

#### Explanation
1. **Command Line Arguments (`os.Args`)**: `os.Args` is a slice of strings containing the arguments passed when running the program. The first argument `os.Args[0]` is always the program path itself.
2. **`os.Open`**: Opens a file for reading. It returns a file pointer that satisfies the `io.Reader` interface.
3. **`file.Read`**: Reads bytes from the file directly into the provided byte slice buffer (`data`).

## Run
```bash
go run main.go <filename>
```
