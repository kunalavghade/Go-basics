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
z