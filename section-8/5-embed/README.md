# 5. Embed

This topic covers the `//go:embed` directive for embedding static files directly into the compiled Go binary.

### Code Example (`main.go`)

```go
package main

import (
	"embed"
	"fmt"
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
```

#### Explanation
1. **`//go:embed` Directive**: This is a special compiler directive. It tells the Go compiler to read the contents of the specified files or directories and embed them into the resulting executable during compilation. 
2. **Embedding into String**: `//go:embed *.txt` over a `string` variable will embed the contents of all `.txt` files directly into that string variable.
3. **Embedding into `embed.FS`**: When embedding entire directories (like `public`), you use `embed.FS`. This creates an in-memory file system. You can then read files from it at runtime using `publc.ReadFile("public/data.txt")` without needing the files to be present on the disk!

## Run
```bash
go run main.go
```
