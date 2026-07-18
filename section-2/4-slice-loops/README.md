# 4. Slice and Loops

This topic covers slices and loop iterations in Go.

### Code Example (`main.go`)

```go
package main

import "fmt"

/*
Arrays vs Slices
- Arrays have a fixed size, while slices are dynamic and can grow or shrink in size.
- Arrays are value types, Slices are reference types.
*/

func main() {
	cards := []string{newCard(), "Two of Hearts", "Three of Diamonds"}
	cards = append(cards, "Four of Clubs")

	// index, curent value := range collection
	for i, card := range cards {
		fmt.Printf("Index: %d, Card: %s\n", i, card)
	}

}

func newCard() string {
	return "Ace of Spades"
}
```

#### Explanation
1. **Slices**: Slices are declared without a size inside the brackets `[]string{...}`. They are dynamic, meaning they can grow or shrink, unlike arrays.
2. **`append`**: The `append(slice, element)` function adds a new element to the end of a slice. It returns a new slice, so we must assign it back to the original variable (`cards = append(cards, ...)`).
3. **Range Loop**: The `range` keyword is used in a `for` loop to iterate over elements. It returns both the `index` (i) and the current `value` (card) of the slice.

## Run
```bash
go run main.go
```
