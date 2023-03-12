package game

import (
	"errors"
	"sort"
)

func calcHighestRankWithCount(cards []Card, exclude RankT, count uint8) (RankT, error) {
	maxRank := NO_RANK
	seen := map[RankT]uint8{}
	for _, card := range cards {
		seen[card.Rank]++
		if seen[card.Rank] >= count &&
			card.Rank > maxRank &&
			card.Rank != exclude {
			maxRank = card.Rank
		}
	}
	if maxRank == NO_RANK {
		return maxRank, errors.New("no rank in cards have sufficient count")
	}
	return maxRank, nil
}

func calcHighestCard(cards []Card) RankT {
	maxRank := NO_RANK
	for _, card := range cards {
		if card.Rank > maxRank {
			maxRank = card.Rank
		}
	}
	return maxRank
}

func calcHighestPair(cards []Card) (RankT, error) {
	maxRank, err := calcHighestRankWithCount(cards, NO_RANK, 2)
	if err != nil {
		return maxRank, errors.New("no pair found")
	}
	return maxRank, nil
}

func calcHighestTwoPair(cards []Card) (RankT, RankT, error) {
	firstRank := NO_RANK
	if rank, err := calcHighestRankWithCount(cards, NO_RANK, 2); err != nil {
		return NO_RANK, NO_RANK, errors.New("no two pair found")
	} else {
		firstRank = rank
	}

	secondRank := NO_RANK
	if rank, err := calcHighestRankWithCount(cards, firstRank, 2); err != nil {
		return NO_RANK, NO_RANK, errors.New("no two pair found")
	} else {
		secondRank = rank
	}

	return firstRank, secondRank, nil
}

func calcHighestThreeOfAKind(cards []Card) (RankT, error) {
	maxRank, err := calcHighestRankWithCount(cards, NO_RANK, 3)
	if err != nil {
		return maxRank, errors.New("no three of a kind found")
	}
	return maxRank, nil
}

func calcHighestStraight(cards []Card) (RankT, error) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	maxRank := NO_RANK
	increasingCount := 0
	if cards[0].Rank == TWO && cards[len(cards)-1].Rank == ACE {
		increasingCount++
	}
	for i, card := range cards[1:] {
		if card.Rank == cards[i].Rank {
			continue
		}
		if card.Rank == cards[i].Rank+1 {
			increasingCount++
		} else {
			increasingCount = 0
		}
		if increasingCount >= 4 {
			maxRank = card.Rank
		}
	}
	if maxRank == NO_RANK {
		return maxRank, errors.New("no straight found")
	}
	return maxRank, nil
}

func calcHighestFlush(cards []Card) (SuitT, RankT, error) {
	maxRank := NO_RANK
	flushSuit := NO_SUIT
	seen := map[SuitT]*struct {
		count uint8
		rank  RankT
	}{}
	for _, card := range cards {
		if card.Rank > seen[card.Suit].rank {
			seen[card.Suit].rank = card.Rank
		}
		seen[card.Suit].count++
		if seen[card.Suit].count >= 5 {
			maxRank = seen[card.Suit].rank
			flushSuit = card.Suit
		}
	}
	if flushSuit == NO_SUIT {
		return flushSuit, maxRank, errors.New("no flush found")
	}
	return flushSuit, maxRank, nil
}

func calcHighestFullHouse(cards []Card) (RankT, RankT, error) {
	firstRank := NO_RANK
	if rank, err := calcHighestRankWithCount(cards, NO_RANK, 3); err != nil {
		return NO_RANK, NO_RANK, errors.New("no full house found")
	} else {
		firstRank = rank
	}

	secondRank := NO_RANK
	if rank, err := calcHighestRankWithCount(cards, firstRank, 2); err != nil {
		return NO_RANK, NO_RANK, errors.New("no full house found")
	} else {
		secondRank = rank
	}

	return firstRank, secondRank, nil
}

func calcHighestFourOfAKind(cards []Card) (RankT, error) {
	maxRank, err := calcHighestRankWithCount(cards, NO_RANK, 4)
	if err != nil {
		return maxRank, errors.New("no four of a kind found")
	}
	return maxRank, nil
}

func calcHighestStraightFlush(cards []Card) (RankT, error) {
	maxRank := NO_RANK
	suitMap := map[SuitT][]Card{}
	for _, card := range cards {
		suitMap[card.Suit] = append(suitMap[card.Suit], card)
	}
	for _, cardSlice := range suitMap {
		if rank, err := calcHighestStraight(cardSlice); err == nil && rank > maxRank {
			maxRank = rank
		}
	}
	if maxRank == NO_RANK {
		return NO_RANK, errors.New("no straight flush found")
	}
	return maxRank, nil
}
