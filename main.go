package main

import (
	"fmt"
	"poker-game/game"
)

func main() {
	deck := game.NewDeck()

	A := &game.Player{
		Username: "P1",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	B := &game.Player{
		Username: "P2",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	C := &game.Player{
		Username: "P3",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	D := &game.Player{
		Username: "P4",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	E := &game.Player{
		Username: "P5",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	F := &game.Player{
		Username: "P6",
		Stack:    100,
		Hand:     []game.Card{deck.Pop(), deck.Pop()},
	}

	table := game.NewTable(1, 2, []*game.Player{A, B, C, D, E, F})

	table.Community = []game.Card{
		deck.Pop(),
		deck.Pop(),
		deck.Pop(),
		deck.Pop(),
		deck.Pop(),
	}

	fmt.Println(table)
}
