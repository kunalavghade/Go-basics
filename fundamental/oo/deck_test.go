package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%s'", d[0])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%s'", loadedDeck[0])
	}

	os.Remove(filename)
}
