package game

import (
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

// Function to evaluate the best hand given a slice of Cards.
// It returns the HandRank and a Card slice representing the winning hand.
// The order of the cards in the slice is the tie-breaking order.
func EvaluateBestHand(cards []Card) (HandRank, []Card) {
	if straightFlush := calcHighestStraightFlush(cards); straightFlush != nil {
		return STRAIGHT_FLUSH, straightFlush
	}

	if fourOfAKind := calcHighestFourOfAKind(cards); fourOfAKind != nil {
		return FOUR_OF_A_KIND, fourOfAKind
	}

	if fullHouse := calcHighestFullHouse(cards); fullHouse != nil {
		return FULL_HOUSE, fullHouse
	}

	if flush := calcHighestFlush(cards); flush != nil {
		return FLUSH, flush
	}

	if straight := calcHighestStraight(cards); straight != nil {
		return STRAIGHT, straight
	}

	if threeOfAKind := calcHighestThreeOfAKind(cards); threeOfAKind != nil {
		return THREE_OF_A_KIND, threeOfAKind
	}

	if twoPair := calcHighestTwoPair(cards); twoPair != nil {
		return TWO_PAIR, twoPair
	}

	if pair := calcHighestPair(cards); pair != nil {
		return PAIR, pair
	}

	return HIGH_CARD, calcHighestCard(cards)
}

func calcCardsOfEachRank(cards []Card) map[Rank][]Card {
	cardsOfEachRank := map[Rank][]Card{}
	for _, card := range cards {
		cardsOfEachRank[card.rank] = append(cardsOfEachRank[card.rank], card)
	}
	return cardsOfEachRank
}

func calcCardsOfEachSuit(cards []Card) map[Suit][]Card {
	cardsOfEachSuit := map[Suit][]Card{}
	for _, card := range cards {
		cardsOfEachSuit[card.suit] = append(cardsOfEachSuit[card.suit], card)
	}
	return cardsOfEachSuit
}

func calcKickers(cards []Card, excludeRank mapset.Set[Rank], count uint8) []Card {
	kickers := []Card{}
	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := ACE; rank >= TWO; rank-- {
		if len(cardsOfEachRank[rank]) > 0 && !excludeRank.Contains(rank) {
			kickers = append(kickers, cardsOfEachRank[rank][0])
			excludeRank.Add(rank)
			count--
		}
		if count == 0 {
			return kickers
		}
	}
	return nil
}

func calcHighestCard(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank > cards[j].rank
	})
	return cards[:5]
}

func calcHighestPair(cards []Card) []Card {
	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := ACE; rank >= TWO; rank-- {
		if len(cardsOfEachRank[rank]) == 2 {
			return append(cardsOfEachRank[rank], calcKickers(cards, mapset.NewSet(rank), 3)...)
		}
	}
	return nil
}

func calcHighestTwoPair(cards []Card) []Card {
	if len(cards) < 4 {
		return nil
	}

	twoPair := []Card{}

	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := ACE; rank >= TWO; rank-- {

		if len(cardsOfEachRank[rank]) == 2 {
			twoPair = append(twoPair, cardsOfEachRank[rank]...)
		}

		if len(twoPair) == 4 {
			return append(twoPair, calcKickers(cards, mapset.NewSet(twoPair[0].rank, twoPair[2].rank), 1)...)
		}
	}

	return nil
}

func calcHighestThreeOfAKind(cards []Card) []Card {
	if len(cards) < 3 {
		return nil
	}

	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := ACE; rank >= TWO; rank-- {
		if len(cardsOfEachRank[rank]) == 3 {
			return append(cardsOfEachRank[rank], calcKickers(cards, mapset.NewSet(rank), 2)...)
		}
	}
	return nil
}

func calcHighestStraight(cards []Card) []Card {
	//If the slice of cards contains less than 5 elements, early return nil
	if len(cards) < 5 {
		return nil
	}

	//Sort the slice in descending order
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank > cards[j].rank
	})

	//Initialize a slice with the first element in cards parameter
	cardsOfStraight := []Card{cards[0]}

	//Iterate through cards parameter excluding the first element
	for i, card := range cards[1:] {
		//Check if current card is 1 less than previous card of slice sorted in descending order
		if card.rank == cards[i].rank-1 {
			cardsOfStraight = append(cardsOfStraight, card)
		} else if card.rank != cards[i].rank {
			cardsOfStraight = []Card{card}
		}

		//Check if cardsOfStraight has 5 cards in it
		if len(cardsOfStraight) == 5 {
			return cardsOfStraight
		}
	}

	//Account for edge case of 5-high straight
	if cards[len(cards)-1].rank == TWO && cards[0].rank == ACE {
		cardsOfStraight = append(cardsOfStraight, cards[0])
	}

	//Check straight condition after accounting for edge case
	if len(cardsOfStraight) == 5 {
		return cardsOfStraight
	}

	return nil
}

func calcHighestFlush(cards []Card) []Card {
	if len(cards) < 5 {
		return nil
	}

	cardsOfEachSuit := calcCardsOfEachSuit(cards)
	for _, cardsOfOneSuit := range cardsOfEachSuit {
		if len(cardsOfOneSuit) >= 5 {
			sort.Slice(cardsOfOneSuit, func(i, j int) bool {
				return cardsOfOneSuit[i].rank > cardsOfOneSuit[j].rank
			})
			return cardsOfOneSuit[:5]
		}
	}
	return nil
}

func calcHighestFullHouse(cards []Card) []Card {
	if len(cards) < 5 {
		return nil
	}

	var threeOfAKind []Card
	var pair []Card

	cardsOfEachRank := calcCardsOfEachRank(cards)

	for rank := ACE; rank >= TWO; rank-- {
		if len(cardsOfEachRank[rank]) == 3 {
			threeOfAKind = cardsOfEachRank[rank]
		} else if len(cardsOfEachRank[rank]) == 2 {
			pair = cardsOfEachRank[rank]
		}

		if threeOfAKind != nil && pair != nil {
			return append(threeOfAKind, pair...)
		}
	}

	return nil
}

func calcHighestFourOfAKind(cards []Card) []Card {
	if len(cards) < 4 {
		return nil
	}

	cardsOfEachRank := calcCardsOfEachRank(cards)

	for rank := ACE; rank >= TWO; rank-- {
		if len(cardsOfEachRank[rank]) == 4 {
			return append(cardsOfEachRank[rank], calcKickers(cards, mapset.NewSet(rank), 1)...)
		}
	}

	return nil
}

func calcHighestStraightFlush(cards []Card) []Card {
	if len(cards) < 5 {
		return nil
	}

	cardsOfEachSuit := calcCardsOfEachSuit(cards)

	for _, cardsOfOneSuit := range cardsOfEachSuit {
		straightFlush := calcHighestStraight(cardsOfOneSuit)
		if straightFlush != nil {
			return straightFlush
		}
	}

	return nil
}
