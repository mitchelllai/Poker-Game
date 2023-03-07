package player

import (
	"fmt"
	"poker-game/game"
)

type Player struct {
	username string
	stack    float64
	hand     []game.Card
}

func (player Player) String() string {
	hand := fmt.Sprint(player.hand)
	stack := fmt.Sprintf("%.2f", player.stack)
	return "Username: " + player.username + "\n" +
		"Stack: $" + stack + "\n" +
		"Hand: " + hand + "\n"
}

func NewPlayer(username string, buyIn float64, hand []game.Card) Player {
	return Player{username, buyIn, hand}
}
