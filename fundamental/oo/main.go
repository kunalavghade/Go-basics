package main

/*
Arrays vs Slices
- Arrays have a fixed size, while slices are dynamic and can grow or shrink in size.
- Arrays are value types, meaning that when you assign an array to another variable, it creates a copy of the array. Slices, on the other hand, are reference types, meaning that when you assign a slice to another variable, both variables point to the same underlying array.
- Arrays are less flexible than slices, as they cannot be resized or easily manipulated. Slices provide more functionality and are generally more convenient to work with in Go.

Loops in Go
- Go provides two types of loops: for loops and range loops.
- For loops are used for iterating over a sequence of values, such as an array or slice. They consist of three components: initialization, condition, and post statement.
- Range loops are used for iterating over elements in a collection, such as an array, slice, map, or string. They provide a convenient way to access both the index and value of each element in the collection.
*/

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	// index, curent value := range collection
	// cards.print()
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}

func newCard() string {
	return "Ace of Spades"
}
