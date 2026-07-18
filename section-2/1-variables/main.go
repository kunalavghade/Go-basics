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
