package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // reuse the code right here
	fmt.Println("b.getGreeting(): ", b.getGreeting())
}

// func printGreeting2(sb spanishBot) {
// 	fmt.Println("sb.getGreeting(): ", sb.getGreeting())
// }

// cannot have functions with the same name 和 C++ 不同，不能重构函数

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// interfaces are not generic types
