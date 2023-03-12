package game

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Deck struct {
	Cards mapset.Set[Card]
}

func (deck Deck) Pop() Card {
	card, _ := deck.Cards.Pop()
	return card
}

func NewDeck() Deck {
	cards := mapset.NewSet[Card]()
	for rank := TWO; rank <= ACE; rank++ {
		for suit := SPADE; suit <= DIAMOND; suit++ {
			cards.Add(Card{rank: rank, suit: suit})
		}
	}
	return Deck{cards}
}
