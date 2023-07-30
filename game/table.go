package game

import (
	"fmt"
	"math/rand"
)

type Table struct {
	Id         int
	SmallBlind float64
	BigBlind   float64
	PlayerMap  map[Position]*Player
	Action     *Player
	Pot        float64
	Community  []Card
}

func (table Table) String() string {
	id := "Table Id: " + fmt.Sprint(table.Id) + "\n\n"
	stakes := "Stakes: " + fmt.Sprint(table.SmallBlind) + "/" + fmt.Sprint(table.BigBlind) + "\n\n"
	community := "Community Cards: " + fmt.Sprint(table.Community) + "\n\n"
	action := "Action on Player " + fmt.Sprint(table.Action.Username) + "\n\n"
	pot := "Pot: " + fmt.Sprint(table.Pot) + "\n\n"

	players := ""
	initialAction := table.Action
	for ok := true; ok; ok = initialAction != table.Action {
		players += fmt.Sprint(*table.Action)
		handRank, _ := CalcBestHand(append(table.Action.Hand, table.Community...))
		players += "Best Hand: " + fmt.Sprint(handRank) + "\n\n"
		table.Action = table.Action.NextPlayer
	}

	winners := "Winners: "
	for _, winner := range EvaluateWinners(table) {
		winners += winner.Username + " "
	}

	return id + stakes + players + community + action + pot + winners
}

func NewTable(smallBlind float64, bigBlind float64, players []*Player) *Table {
	table := &Table{}
	table.Id = rand.Intn(100)
	table.SmallBlind = smallBlind
	table.BigBlind = bigBlind
	table.PlayerMap = map[Position]*Player{}
	playerCount := len(players)

	for i, player := range players {
		position := Position(i + 1)
		player.Position = position
		table.PlayerMap[position] = player

		if i == 0 {
			player.NextPlayer = players[playerCount-1]
		} else {
			player.NextPlayer = players[i-1]
		}

	}

	table.Action = players[playerCount-1]

	return table
}

func EvaluateWinners(table Table) []*Player {
	winners := []*Player{table.Action}
	bestHandRank, bestHand := CalcBestHand(append(table.Action.Hand, table.Community...))

	initialAction := table.Action
	table.Action = table.Action.NextPlayer
	for ok := true; ok; ok = initialAction != table.Action {
		playerHandRank, playerHand := CalcBestHand(append(table.Action.Hand, table.Community...))

		if playerHandRank > bestHandRank {
			winners = []*Player{table.Action}
			bestHandRank = playerHandRank
			bestHand = playerHand
		} else if playerHandRank == bestHandRank {
			for i, bestHandCard := range bestHand {

				if playerHand[i].Rank > bestHandCard.Rank {
					winners = []*Player{table.Action}
					bestHand = playerHand
					break
				}

				if playerHand[i].Rank < bestHandCard.Rank {
					break
				}

				if i == len(bestHand)-1 {
					winners = append(winners, table.Action)
				}
			}
		}

		table.Action = table.Action.NextPlayer
	}

	return winners
}
