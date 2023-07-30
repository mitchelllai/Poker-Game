package game

import (
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

// Calculate the best hand given a slice of Cards.
// Returns the HandRank and a Card slice representing the winning hand.
// The order of the cards in the slice is the tie-breaking order.
func CalcBestHand(cards []Card) (HandRank, []Card) {
	if straightFlush := calcHighestStraightFlush(cards); straightFlush != nil {
		return StraightFlush, straightFlush
	}

	if fourOfAKind := calcHighestFourOfAKind(cards); fourOfAKind != nil {
		return FourOfAKind, fourOfAKind
	}

	if fullHouse := calcHighestFullHouse(cards); fullHouse != nil {
		return FullHouse, fullHouse
	}

	if flush := calcHighestFlush(cards); flush != nil {
		return Flush, flush
	}

	if straight := calcHighestStraight(cards); straight != nil {
		return Straight, straight
	}

	if threeOfAKind := calcHighestThreeOfAKind(cards); threeOfAKind != nil {
		return ThreeOfAKind, threeOfAKind
	}

	if twoPair := calcHighestTwoPair(cards); twoPair != nil {
		return TwoPair, twoPair
	}

	if pair := calcHighestPair(cards); pair != nil {
		return Pair, pair
	}

	return HighCard, calcHighestCard(cards)
}

// Returns a map
// where the key is each rank in cards parameter
// and each value is a slice of cards containing each card of that rank.
func calcCardsOfEachRank(cards []Card) map[Rank][]Card {
	cardsOfEachRank := map[Rank][]Card{}

	for _, card := range cards {
		cardsOfEachRank[card.Rank] = append(cardsOfEachRank[card.Rank], card)
	}

	return cardsOfEachRank
}

// Returns a map
// where the key is each suit in cards parameter
// and each value is a slice of cards containing each card of that suit.
func calcCardsOfEachSuit(cards []Card) map[Suit][]Card {
	cardsOfEachSuit := map[Suit][]Card{}

	for _, card := range cards {
		cardsOfEachSuit[card.Suit] = append(cardsOfEachSuit[card.Suit], card)
	}

	return cardsOfEachSuit
}

// Returns a slice
// where each element is a card in cards parameter
// such that the card's rank is not in excludeRank,
// and count is the length of the returned kickers slice.
func calcKickers(cards []Card, excludeRanks mapset.Set[Rank], kickersLen int) []Card {
	kickers := []Card{}
	cardsOfEachRank := calcCardsOfEachRank(cards)

	//Iterate through each rank in descending order.
	for rank := Ace; rank >= Two; rank-- {

		//If that rank is in cardsOfEachRank
		//and the rank is not in excludeRanks,
		//then add a card from the slice to kickers slice.
		if cardsOfEachRank[rank] != nil && !excludeRanks.Contains(rank) {
			kickers = append(kickers, cardsOfEachRank[rank][0])
			excludeRanks.Add(rank)
		}

		//Return early if kickers reaches length specified in parameter
		if len(kickers) == kickersLen {
			return kickers
		}

	}

	return nil
}

// Returns a slice of 5 cards
// where the cards are sorted in tie-breaking order
func calcHighestCard(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})
	return cards[:5]
}

func calcHighestPair(cards []Card) []Card {
	cardsOfEachRank := calcCardsOfEachRank(cards)

	for rank := Ace; rank >= Two; rank-- {
		if len(cardsOfEachRank[rank]) == 2 {
			excludeRanks := mapset.NewSet(rank)
			kickers := calcKickers(cards, excludeRanks, 3)
			return append(cardsOfEachRank[rank], kickers...)
		}
	}

	return nil
}

func calcHighestTwoPair(cards []Card) []Card {
	if len(cards) < 4 {
		return nil
	}

	var twoPair []Card

	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := Ace; rank >= Two; rank-- {

		if len(cardsOfEachRank[rank]) == 2 {
			twoPair = append(twoPair, cardsOfEachRank[rank]...)
		}

		if len(twoPair) == 4 {
			excludeRanks := mapset.NewSet(twoPair[0].Rank, twoPair[2].Rank)
			return append(twoPair, calcKickers(cards, excludeRanks, 1)...)
		}
	}

	return nil
}

func calcHighestThreeOfAKind(cards []Card) []Card {
	if len(cards) < 3 {
		return nil
	}

	cardsOfEachRank := calcCardsOfEachRank(cards)
	for rank := Ace; rank >= Two; rank-- {
		if len(cardsOfEachRank[rank]) == 3 {
			kickers := calcKickers(cards, mapset.NewSet(rank), 2)
			return append(cardsOfEachRank[rank], kickers...)
		}
	}
	return nil
}

func calcHighestStraight(cards []Card) []Card {
	//If the slice of cards contains less than 5 elements, early return nil
	if len(cards) < 5 {
		return nil
	}

	var cardsOfStraight []Card

	cardsOfEachRank := calcCardsOfEachRank(cards)

	for rank := Ace; rank >= Two; rank-- {
		if cardsOfEachRank[rank] != nil {
			if cardsOfStraight == nil ||
				rank == cardsOfStraight[len(cardsOfStraight)-1].Rank-1 {
				cardsOfStraight = append(cardsOfStraight, cardsOfEachRank[rank][0])
			}

			if len(cardsOfStraight) == 5 {
				return cardsOfStraight
			}

			if rank == Two && cardsOfEachRank[Ace] != nil {
				cardsOfStraight = append(
					cardsOfStraight,
					cardsOfEachRank[Ace][0],
				)
			}

			if len(cardsOfStraight) == 5 {
				return cardsOfStraight
			}
		} else {
			cardsOfStraight = nil
		}

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
				return cardsOfOneSuit[i].Rank > cardsOfOneSuit[j].Rank
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

	for rank := Ace; rank >= Two; rank-- {
		if len(cardsOfEachRank[rank]) == 3 {
			threeOfAKind = cardsOfEachRank[rank]
		}

		if len(cardsOfEachRank[rank]) == 2 {
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

	for rank := Ace; rank >= Two; rank-- {
		if len(cardsOfEachRank[rank]) == 4 {
			kickers := calcKickers(cards, mapset.NewSet(rank), 1)
			return append(cardsOfEachRank[rank], kickers...)
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
