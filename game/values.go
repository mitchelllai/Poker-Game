package game

import "fmt"

type Rank uint8

const (
	NO_RANK Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

func (rank Rank) String() string {
	switch rank {
	case 0:
		return ""
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	case ACE:
		return "A"
	default:
		return fmt.Sprint(uint8(rank) + 1)
	}
}

type Suit uint8

const (
	NO_SUIT Suit = iota
	SPADE
	HEART
	CLUB
	DIAMOND
)

func (suit Suit) String() string {
	switch suit {
	case SPADE:
		return "♠"
	case HEART:
		return "♥"
	case CLUB:
		return "♣"
	case DIAMOND:
		return "♦"
	default:
		return ""
	}
}

type Position uint8

const (
	BB Position = iota + 1
	SB
	BTN
	CO
	HJ
	UTG
)

func (position Position) String() string {
	switch position {
	case BB:
		return "BB"
	case SB:
		return "SB"
	case BTN:
		return "BTN"
	case CO:
		return "CO"
	case HJ:
		return "HJ"
	case UTG:
		return "UTG"
	default:
		return ""
	}
}

type HandRank uint8

const (
	HIGH_CARD HandRank = iota
	PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

const MAX_PLAYER_COUNT = 6
