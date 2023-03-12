package game

import (
	"sort"
)

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
	seen := map[Suit]*struct {
		count uint8
		rank  Rank
	}{}
	for _, card := range cards {
		if card.rank > seen[card.suit].rank {
			seen[card.suit].rank = card.rank
		}
		seen[card.suit].count++
		if seen[card.suit].count >= 5 {
			maxRank = seen[card.suit].rank
		}
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
