# 1. Interfaces

This topic covers the basics of Interfaces in Go.

### Code Example (`main.go`)

```go
package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "¡Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
```

#### Explanation
1. **Implicit Interfaces**: In Go, interfaces are implemented implicitly. Since both `englishBot` and `spanishBot` have a function `getGreeting() string`, they automatically implement the `bot` interface.
2. **Polymorphism**: The `printGreeting` function accepts *any* type that satisfies the `bot` interface. This allows us to pass both `englishBot` and `spanishBot` into the same function.

## Run
```bash
go run main.go
```
