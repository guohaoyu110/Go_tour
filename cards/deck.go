package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// []string(d) =
	// have the ability to turn our deck to a string
}

// ["red", "yellow", "blue"]

func (d deck) saveToFile(filename string) error {
	// return type error
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// use default permissions

}

func newDeckFromFile(filename string) deck {
	// the return type is deck
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// comes to the very particular error
		// option #1 - log the error and return a call to newDeck()
		// option #2 - extremely wrong with our program, and entirely quit the program
		fmt.Println(err)
		os.Exit(1) // vscode automatically generate the package "os"

	}
	s := strings.Split(string(bs), ",")
	// print out the value right now
	// Ace of Spades, two of spades, three of spades
	return deck(s)

}

func (d deck) shuffle() {
	// nothing else to return
	// this is truly random
	source := rand.NewSource(time.Now().UnixNano())
	// get the source and create the new rand object
	r := rand.New(source)
	for i := range d { // do not add cards
		newPositon := r.Intn(len(d) - 1) // the random generator, depends on the seed value
		d[i], d[newPositon] = d[newPositon], d[i]

	}
}

// know how to use the documentation
