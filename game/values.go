package game

import "fmt"

type Rank uint8

const (
	NoRank Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (rank Rank) String() string {
	switch rank {
	case NoRank:
		return ""
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return fmt.Sprint(uint8(rank) + 1)
	}
}

type Suit uint8

const (
	NoSuit Suit = iota
	Spade
	Heart
	Club
	Diamond
)

func (suit Suit) String() string {
	switch suit {
	case Spade:
		return "♠"
	case Heart:
		return "♥"
	case Club:
		return "♣"
	case Diamond:
		return "♦"
	default:
		return ""
	}
}

type Position uint8

const (
	Bb Position = iota + 1
	Sb
	Btn
	Co
	Hj
	Utg
)

func (position Position) String() string {
	switch position {
	case Bb:
		return "BB"
	case Sb:
		return "SB"
	case Btn:
		return "BTN"
	case Co:
		return "CO"
	case Hj:
		return "HJ"
	case Utg:
		return "UTG"
	default:
		return ""
	}
}

type HandRank uint8

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (handRank HandRank) String() string {
	switch handRank {
	case HighCard:
		return "High Card"
	case Pair:
		return "Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case StraightFlush:
		return "Straight Flush"
	default:
		return ""
	}

}
