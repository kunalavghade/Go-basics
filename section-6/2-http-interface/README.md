# 2. HTTP Interface

This topic covers implementing HTTP interfaces and custom IO writers.

### Code Example (`main.go`)

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	io.Copy(logWriter{}, resp.Body)
}
```

#### Explanation
1. **`http.Get`**: The `net/http` package allows us to make network requests. The response body implements the `io.Reader` interface.
2. **Custom `io.Writer`**: We created a custom struct `logWriter` and attached a `Write` method to it. By doing so, it now satisfies the `io.Writer` interface. 
3. **`io.Copy`**: We use `io.Copy` to stream the data from the `resp.Body` (a Reader) directly into our `logWriter` (a Writer).

## Run
```bash
go run main.go
```
