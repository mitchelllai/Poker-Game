package game

import (
	"fmt"
)

type Card struct {
	Rank uint8
	Suit uint8
}

func NewCard(rank uint8, suit uint8) Card {
	return Card{rank, suit}
}

func (card Card) String() string {
	var rank, suit string

	switch card.Rank {
	case 0:
		rank = ""
	case JACK:
		rank = "J"
	case QUEEN:
		rank = "Q"
	case KING:
		rank = "K"
	case ACE:
		rank = "A"
	default:
		rank = fmt.Sprint(card.Rank + 1)
	}

	switch card.Suit {
	case 0:
		suit = ""
	case SPADE:
		suit = "♠"
	case HEART:
		suit = "♥"
	case CLUB:
		suit = "♣"
	case DIAMOND:
		suit = "♦"
	}

	return rank + suit
}
