# 5. Object Oriented

This topic covers Object-Oriented programming style in Go using custom types and receivers.

### Code Example (`main.go`)

```go
package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	// cards.print()
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
```

*(Note: The actual implementations for `newDeck`, `deal`, `print`, `saveToFile`, and `shuffle` are located in `deck.go` in the same directory, acting as receiver functions on a custom `deck` type.)*

#### Explanation
1. **Custom Types**: Go doesn't have classes. Instead, we can create custom types based on existing ones (e.g., `type deck []string`) or `structs`.
2. **Receiver Functions**: Methods like `cards.print()` or `cards.shuffle()` are made possible by writing functions with a **receiver**. A receiver binds a function to a specific type, making it act like a method on a class instance.
3. **Multiple Return Values**: Functions in Go can return multiple values. For instance, `deal(cards, 5)` returns two variables: `hand` and `remainingCards`.

## Run
```bash
go run main.go deck.go
```
