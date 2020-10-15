package main

import "fmt"

// Create a new type of 'deck'
// type deck is a slice of strings
type deck [] string

func newDeck() deck {
	cards := deck {}

	cardSuits := [] string {
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}

	cardValues := [] string {
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}