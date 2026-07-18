# 3. Temporary Files

This topic covers creating and managing temporary files and directories.

### Code Example (`main.go`)

```go
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
```

#### Explanation
1. **`os.CreateTemp`**: Creates a new temporary file in the default directory for temporary files (first argument `""`). The second argument provides a pattern for the file name.
2. **`os.MkdirTemp`**: Similar to `CreateTemp`, but creates a temporary directory. The `*` in the pattern is replaced by a random string to guarantee a unique name.
3. **Deferred Cleanup**: We use `defer os.Remove(tmpfile.Name())` and `defer os.RemoveAll(tmpDir)` to guarantee that these temporary files and folders are cleaned up before the program exits, preventing disk clutter.

## Run
```bash
go run main.go
```
