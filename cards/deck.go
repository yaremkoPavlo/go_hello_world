package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Cluds"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eigth",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}
	for _, suit := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bt, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bt), ",")
	return deck(s)
}

func (d deck) shuffle() {
	sourse := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sourse)

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}
