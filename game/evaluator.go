package game

import (
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

// Function to evaluate the best hand given a slice of Cards.
// It returns the HandRank and a slice of Ranks used as tie-breakers.
func EvaluateBestHand(cards []Card) (HandRank, []Rank) {
	if rank := calcHighestStraightFlush(cards); rank != NO_RANK {
		return STRAIGHT_FLUSH, []Rank{rank}
	}
	if rank := calcHighestFourOfAKind(cards); rank != NO_RANK {
		return FOUR_OF_A_KIND, append([]Rank{rank}, calcKickers(cards, mapset.NewSet(rank), 1)...)
	}
	if firstRank, secondRank := calcHighestFullHouse(cards); firstRank != NO_RANK {
		return FULL_HOUSE, []Rank{firstRank, secondRank}
	}
	if rank := calcHighestFlush(cards); rank != NO_RANK {
		return FLUSH, []Rank{rank}
	}
	if rank := calcHighestStraight(cards); rank != NO_RANK {
		return STRAIGHT, []Rank{rank}
	}
	if rank := calcHighestThreeOfAKind(cards); rank != NO_RANK {
		return THREE_OF_A_KIND, append([]Rank{rank}, calcKickers(cards, mapset.NewSet(rank), 2)...)
	}
	if firstRank, secondRank := calcHighestTwoPair(cards); firstRank != NO_RANK {
		return TWO_PAIR, append([]Rank{firstRank, secondRank}, calcKickers(cards, mapset.NewSet(firstRank, secondRank), 1)...)
	}
	if rank := calcHighestPair(cards); rank != NO_RANK {
		return PAIR, append([]Rank{rank}, calcKickers(cards, mapset.NewSet(rank), 3)...)
	}

	return HIGH_CARD, calcKickers(cards, mapset.NewSet[Rank](), 5)
}

func calcKickers(cards []Card, exclude mapset.Set[Rank], count uint8) []Rank {
	kickers := []Rank{}
	prev := NO_RANK
	for i := uint8(0); i < count; i++ {
		if prev != NO_RANK {
			exclude.Add(prev)
		}
		prev = calcHighestRankWithCount(cards, exclude, 1)
		kickers = append(kickers, prev)
	}
	return kickers
}

// Function to calculate the highest Rank given, a slice of Cards, a set of Ranks to exclude, and the count for the Rank.
// Returns the highest Rank with count.
func calcHighestRankWithCount(cards []Card, exclude mapset.Set[Rank], count uint8) Rank {
	maxRank := NO_RANK
	seen := map[Rank]uint8{}
	for _, card := range cards {
		seen[card.rank]++
		if seen[card.rank] >= count &&
			card.rank > maxRank &&
			!exclude.Contains(card.rank) {
			maxRank = card.rank
		}
	}
	return maxRank
}

// Function to calculate the high card given a slice of Cards.
// Returns the Rank of the high card.
func calcHighestCard(cards []Card) Rank {
	maxRank := NO_RANK
	for _, card := range cards {
		if card.rank > maxRank {
			maxRank = card.rank
		}
	}
	return maxRank
}

// Function to calculate the highest pair given a slice of Cards.
// Returns the Rank of the highest pair.
func calcHighestPair(cards []Card) Rank {
	exclude := mapset.NewSet[Rank]()
	return calcHighestRankWithCount(cards, exclude, 2)
}

// Function to calculate the highest two pair given a slice of Cards.
// Returns the Ranks of both pairs.
func calcHighestTwoPair(cards []Card) (Rank, Rank) {
	//Define an empty exclude set.
	exclude := mapset.NewSet[Rank]()

	firstRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, exclude, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		firstRank = rank
	}

	exclude.Add(firstRank)

	secondRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, exclude, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		secondRank = rank
	}

	return firstRank, secondRank
}

func calcHighestThreeOfAKind(cards []Card) Rank {
	exclude := mapset.NewSet[Rank]()
	return calcHighestRankWithCount(cards, exclude, 3)
}

func calcHighestStraight(cards []Card) Rank {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank < cards[j].rank
	})
	maxRank := NO_RANK
	cardsLen := len(cards) - 1
	increasingCount := 0
	if cards[0].rank == TWO && cards[cardsLen].rank == ACE {
		increasingCount++
	}
	for i, card := range cards[1:] {
		if card.rank == cards[i].rank {
			continue
		}
		if card.rank == cards[i].rank+1 {
			increasingCount++
		} else {
			increasingCount = 0
		}
		if increasingCount >= 4 {
			maxRank = card.rank
		}
	}
	return maxRank
}

func calcHighestFlush(cards []Card) Rank {
	maxRank := NO_RANK
	seen := map[Suit]struct {
		count uint8
		rank  Rank
	}{}
	for _, card := range cards {
		count, rank := seen[card.suit].count, seen[card.suit].rank
		if card.rank > rank {
			rank = card.rank
		}
		count++
		if count >= 5 {
			maxRank = rank
		}
		seen[card.suit] = struct {
			count uint8
			rank  Rank
		}{count, rank}
	}
	return maxRank
}

func calcHighestFullHouse(cards []Card) (Rank, Rank) {
	exclude := mapset.NewSet[Rank]()

	firstRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, exclude, 3); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		firstRank = rank
	}

	exclude.Add(firstRank)

	secondRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, exclude, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		secondRank = rank
	}

	return firstRank, secondRank
}

func calcHighestFourOfAKind(cards []Card) Rank {
	exclude := mapset.NewSet[Rank]()
	return calcHighestRankWithCount(cards, exclude, 4)
}

func calcHighestStraightFlush(cards []Card) Rank {
	maxRank := NO_RANK
	suitMap := map[Suit][]Card{}
	for _, card := range cards {
		suitMap[card.suit] = append(suitMap[card.suit], card)
	}
	for _, cardSlice := range suitMap {
		if rank := calcHighestStraight(cardSlice); rank > maxRank {
			maxRank = rank
		}
	}
	return maxRank
}
