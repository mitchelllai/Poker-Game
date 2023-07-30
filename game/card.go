package game

import "fmt"

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return fmt.Sprint(card.Rank) + fmt.Sprint(card.Suit)
}
