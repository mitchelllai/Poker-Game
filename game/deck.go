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
	for rank := Two; rank <= Ace; rank++ {
		for suit := Spade; suit <= Diamond; suit++ {
			cards.Add(Card{Rank: rank, Suit: suit})
		}
	}
	return Deck{cards}
}
