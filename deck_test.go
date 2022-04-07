package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %s", deck[len(deck)-1])
	}
}
