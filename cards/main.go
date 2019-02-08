package main

import "fmt"

func main() {
	//cards := newDeck()
	// hand, renainingCards := deal(cards, 5)
	// hand.print()
	// renainingCards.print()
	//str := cards.toString()
	//fmt.Println(str)
	//cards.saveToFile("my_cards")
	fmt.Println(newDeckFromFile("my_cards"))
}
