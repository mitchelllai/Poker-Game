package main

import (
	"fmt"
	"poker-game/game"
	"poker-game/player"
)

func main() {
	deck := game.NewDeck()
	card1, _ := deck.Pop()
	card2, _ := deck.Pop()
	startingHand := []game.Card{}
	fmt.Println(startingHand)
	startingHand = append(startingHand, card1)
	startingHand = append(startingHand, card2)
	fmt.Println(
		player.NewPlayer(
			"mitchell-lai",
			999.9998708489189,
			startingHand,
		),
	)
	card, _ := deck.Pop()
	fmt.Println(card)
	card, _ = deck.Pop()
	fmt.Println(card)
	card, _ = deck.Pop()
	fmt.Println(card)
	card, _ = deck.Pop()
	fmt.Println(card)
	card, _ = deck.Pop()
	fmt.Println(card)
}
