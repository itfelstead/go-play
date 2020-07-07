package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("wrong 1st card %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("wrong last card %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("loaded wrnong number, expected 16 %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
