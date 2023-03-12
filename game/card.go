package game

import "fmt"

type Card struct {
	rank Rank
	suit Suit
}

func (card Card) String() string {
	return fmt.Sprint(card.rank) + fmt.Sprint(card.suit)
}
