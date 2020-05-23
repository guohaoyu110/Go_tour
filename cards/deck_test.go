package main

import (
	"os"
	"testing"
)

// run file tests is to test all the files in this folder

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 { //一共有16种不同的组合
		t.Errorf("Expected deck length of 16, but got %v", len(d))

	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])

	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// write a test to make sure we can make a deck, save it and load it back up
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 { // if change to 160, it will definitely fail
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
