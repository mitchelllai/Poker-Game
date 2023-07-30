package game

import (
	"testing"
)

func TestCalcHighestCard(t *testing.T) {
	cards := []Card{
		{Ace, Spade},
		{Three, Diamond},
		{Five, Club},
		{Seven, Heart},
		{Nine, Spade},
		{Jack, Diamond},
		{King, Club},
	}
	expectedBestHandRank := HighCard
	expectedBestHand := []Card{
		{Ace, Spade},
		{King, Club},
		{Jack, Diamond},
		{Nine, Spade},
		{Seven, Heart},
	}
	bestHandRank, bestHand := CalcBestHand(cards)
	if bestHandRank != expectedBestHandRank {
		t.Errorf(
			"Result was incorrect; expected best hand rank: %s; actual best hand rank: %s;",
			expectedBestHandRank,
			bestHandRank,
		)
	}
	for i, card := range bestHand {
		if card != expectedBestHand[i] {
			t.Errorf(
				"Result was incorrect; expected best hand: %s; actual best hand: %s",
				expectedBestHand,
				bestHand,
			)
		}
	}
}

func TestCalcHighestPair(t *testing.T) {

}

func TestCalcHighestTwoPair(t *testing.T) {

}

func TestCalcHighestThreeOfAKind(t *testing.T) {

}

func TestCalcHighestStraight(t *testing.T) {

}

func TestCalcHighestFlush(t *testing.T) {

}

func TestCalcHighestFullHouse(t *testing.T) {

}

func TestCalcHighestFourOfAKind(t *testing.T) {

}

func TestCalcHighestStraightFlush(t *testing.T) {

}
