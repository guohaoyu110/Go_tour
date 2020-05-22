package main

import "fmt"

// crate a new type of deck
// which is a slice of strings
// go run main.go deck.go
type deck []string

// 定义一个print的function
// by convention, we just call it d
func (myDeck deck) print() {
	for i, card := range myDeck {
		fmt.Println(i, card)
	}
}
