# 2. Functions

This topic covers functions and return types in Go.

### Code Example (`main.go`)

```go
package main

import "fmt"

func main() {
	card := "Ace of Spades"
	fmt.Println(card)

	card = "Queen of Hearts"
	fmt.Println(card)

	card = newCard()
	fmt.Println(card)
}

// newCard - name of func  |  string - return type of func
func newCard() string {
	return "Five of Diamonds"
}
```

#### Explanation
1. **Variable Reassignment**: Variables initialized with `:=` can be reassigned later using `=` (e.g., `card = "Queen of Hearts"`). You only use `:=` for the initial declaration.
2. **Function Declaration**: Custom functions are declared using `func functionName() returnType`. In this example, `func newCard() string` means `newCard` is a function that returns a `string`.
3. **Return Statements**: Functions return values using the `return` keyword, which must match the declared return type.

## Run
```bash
go run main.go
```
