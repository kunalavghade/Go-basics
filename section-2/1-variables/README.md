# 1. Variables

This topic covers variable declaration and initialization in Go.

### Code Example (`main.go`)

```go
package main

import "fmt"

/*
Basic Go Types:
- string | "Hello, World!"
- int | 42 , -1, 0
- float64 | 3.14, -0.001
- bool | true, false
*/

func main() {
	// var - new var  |  card - name of var  |  string - type of var  |  "Ace of Spades" - value of var
	var card string = "Ace of Spades"
	fmt.Println(card)

	// new way to declare var
	card2 := "Queen of Hearts"
	fmt.Println(card2)
}
```

#### Explanation
1. **Explicit Declaration**: Using `var name type = value` explicitly declares a variable with a specific type (e.g., `var card string`).
2. **Short Variable Declaration**: Using `:=` is a shorthand for declaring and initializing a variable. Go infers the type automatically based on the assigned value (e.g., `card2 := "Queen of Hearts"` infers `string`). This can only be used inside a function.
3. **Basic Types**: Go includes basic types like `string`, `int`, `float64`, and `bool`.

## Run
```bash
go run main.go
```
