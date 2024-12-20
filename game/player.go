package game

import (
	"fmt"
)

type Player struct {
	Username   string
	Stack      float32
	Hand       []Card
	Position   Position
	NextPlayer *Player
}

func (player Player) String() string {
	hand := fmt.Sprint(player.Hand)
	stack := fmt.Sprintf("%.2f", player.Stack)
	return "Username: " + player.Username + "\n" +
		"Stack: $" + stack + "\n" +
		"Position: " + fmt.Sprint(player.Position) + "\n" +
		"Hand: " + hand + "\n"
}
