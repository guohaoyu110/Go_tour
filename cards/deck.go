package main

import (
	"fmt"
	"strings"
)

// crate a new type of deck
// which is a slice of strings
// go run main.go deck.go
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // make this a slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits { // 用_代替，代表我知道这里有一个变量，但是我不想用了
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	// return an actual deck
	return cards
}

// 定义一个print的function
// by convention, we just call it d
func (myDeck deck) print() {
	for i, card := range myDeck {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// annotate the return types 一定要标注返回类型
	fmt.Println(d)
	return d[:handSize], d[handSize:]
}

// func toByteSlice() {

// }

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// []string(d) =
}

// ["red", "yellow", "blue"]
