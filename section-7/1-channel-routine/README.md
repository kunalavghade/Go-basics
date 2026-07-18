# 1. Channel and Routine

This topic covers Goroutines and Channels for concurrency in Go.

### Code Example (`main.go`)

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Link is not working:", link)
		c <- link
		return
	}
	defer resp.Body.Close()
	fmt.Println("Link is working:", link)
	c <- link
}
```

#### Explanation
1. **Goroutines**: The `go` keyword starts a new goroutine (a lightweight thread managed by the Go runtime). `go checkLink(...)` executes the function concurrently.
2. **Channels**: `c := make(chan string)` creates a channel of strings. Channels are used to communicate safely between goroutines.
3. **Channel Communication**: The syntax `c <- link` sends a value into the channel. The syntax `<-c` receives a value from the channel. The `for l := range c` loop continuously receives values from the channel.
4. **Function Literals (Anonymous Functions)**: `go func(link string) { ... }(l)` is an anonymous function launched as a goroutine. It is used here to pause execution (`time.Sleep`) without blocking the main thread before checking the link again.

## Run
```bash
go run main.go
```
