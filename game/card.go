package game

import "fmt"

type Card struct {
	Rank RankT
	Suit SuitT
}

// func NewCard(rank RankT, suit SuitT) Card {
// 	return Card{rank, suit}
// }

func (card Card) String() string {
	return fmt.Sprint(card.Rank) + fmt.Sprint(card.Suit)
}
