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
	players := ""
	for _, playerPtr := range table.PlayerMap {
		players += fmt.Sprint(*playerPtr)
		handRank, _ := EvaluateBestHand(append(playerPtr.Hand, table.Community...))
		players += "Best Hand: " + fmt.Sprint(handRank) + "\n\n"
	}
	community := "Community Cards: " + fmt.Sprint(table.Community) + "\n\n"
	action := "Action on Player " + fmt.Sprint(table.Action.Username) + "\n\n"
	pot := "Pot: " + fmt.Sprint(table.Pot) + "\n\n"
	winner := "Winner: " + fmt.Sprint(EvaluateWinner(table).Username) + "\n"
	return id + stakes + players + community + action + pot + winner
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

func EvaluateWinner(table Table) *Player {
	winner := table.Action
	for _, player := range table.PlayerMap {
		winnerBestHand, winnerTieBreaker := EvaluateBestHand(append(winner.Hand, table.Community...))
		playerBestHand, playerTieBreaker := EvaluateBestHand(append(player.Hand, table.Community...))
		if playerBestHand > winnerBestHand {
			winner = player
		} else if playerBestHand == winnerBestHand {
			for i := 0; i < len(winnerTieBreaker); i++ {
				if playerTieBreaker[i] > winnerTieBreaker[i] {
					winner = player
					break
				}
			}

		}
	}
	return winner
}
