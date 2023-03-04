package game

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Deck mapset.Set[Card]

func NewDeck() Deck {
	deck := mapset.NewSet[Card]()
	for rank := TWO; rank <= ACE; rank++ {
		for suit := SPADE; suit <= DIAMOND; suit++ {
			deck.Add(Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}
