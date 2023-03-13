package game

import (
	"sort"
)

func EvaluateBestHand(cards []Card) (HandRank, []Rank) {
	if rank := calcHighestStraightFlush(cards); rank != NO_RANK {
		return STRAIGHT_FLUSH, []Rank{rank}
	}
	if rank := calcHighestFourOfAKind(cards); rank != NO_RANK {
		return FOUR_OF_A_KIND, []Rank{rank}
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
		return THREE_OF_A_KIND, []Rank{rank}
	}
	if firstRank, secondRank := calcHighestTwoPair(cards); firstRank != NO_RANK {
		return TWO_PAIR, []Rank{firstRank, secondRank}
	}
	if rank := calcHighestPair(cards); rank != NO_RANK {
		return PAIR, []Rank{rank}
	}
	return HIGH_CARD, []Rank{calcHighestCard(cards)}
}

func calcHighestRankWithCount(cards []Card, exclude Rank, count uint8) Rank {
	maxRank := NO_RANK
	seen := map[Rank]uint8{}
	for _, card := range cards {
		seen[card.rank]++
		if seen[card.rank] >= count &&
			card.rank > maxRank &&
			card.rank != exclude {
			maxRank = card.rank
		}
	}
	return maxRank
}

func calcHighestCard(cards []Card) Rank {
	maxRank := NO_RANK
	for _, card := range cards {
		if card.rank > maxRank {
			maxRank = card.rank
		}
	}
	return maxRank
}

func calcHighestPair(cards []Card) Rank {
	return calcHighestRankWithCount(cards, NO_RANK, 2)
}

func calcHighestTwoPair(cards []Card) (Rank, Rank) {
	firstRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, NO_RANK, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		firstRank = rank
	}

	secondRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, firstRank, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		secondRank = rank
	}

	return firstRank, secondRank
}

func calcHighestThreeOfAKind(cards []Card) Rank {
	return calcHighestRankWithCount(cards, NO_RANK, 3)
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
	firstRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, NO_RANK, 3); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		firstRank = rank
	}

	secondRank := NO_RANK
	if rank := calcHighestRankWithCount(cards, firstRank, 2); rank == NO_RANK {
		return NO_RANK, NO_RANK
	} else {
		secondRank = rank
	}

	return firstRank, secondRank
}

func calcHighestFourOfAKind(cards []Card) Rank {
	return calcHighestRankWithCount(cards, NO_RANK, 4)
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
