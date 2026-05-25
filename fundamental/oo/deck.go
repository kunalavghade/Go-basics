package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create new deck type
type deck []string

// receiver function to print the deck
/*
Syntax: func (receiverName receiverType) functionName(parameters) returnType {
	// function body
}
*/
func (d deck) print() {
	// print the index and value of each card in the deck
	for i, card := range d {
		fmt.Printf("Index: %d, Card: %s\n", i, card)
	}
}

func newDeck() deck {
	// create a new deck of cards
	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsSuits {
		for _, values := range cardsValues {
			cards = append(cards, values+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// return two decks, one with the hand size and the other with the remaining cards
	/*
		Syntax. of slicing: slice[low:high] or slice[low:] or slice[:high]
		- [fromIdxinclusive:toIdxexclusive]
		- low: the index of the first element to include in the slice (inclusive)
		- high: the index of the first element to exclude from the slice (exclusive)
		- if low is omitted, it defaults to 0
		- if high is omitted, it defaults to the length of the slice
	*/
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// convert the deck to a string4
	// syntax: strings.Join([]string(slice), separator)
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// save the deck to a file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// read the deck from a file
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	// shuffle the deck
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
