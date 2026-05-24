package main

import "fmt"

// create new deck type
type deck []string

// receiver function to print the deck
/*
Syntax: func (receiverName receiverType) functionName(parameters) returnType {
	// function body
}
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Printf("Index: %d, Card: %s\n", i, card)
	}
}
