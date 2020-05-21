package main

import "fmt"

func main() {
	// the warning is totally fine
	// printState()                      // 必须是go run main.go state.go这一行才能执行
	var card string = "Ace of Spades" // one way to define a variable
	card1 := "Ace1 of Spades"         // 只有第一次定义的时候需要加上冒号
	card = "Five of Diomands"
	fmt.Println(card)
	fmt.Println(card1)
	fmt.Println("card: ", card) // add snippet to better utilize vscode

	card2 := newCard()
	fmt.Println("card2: ", card2)

	fmt.Println(newCard())

}

// a separate function besides main
func newCard() string { // update function declaration, return a string
	return "Five of Diamonds and Ace of Spades"
}

// read the error message and see what it gives us
