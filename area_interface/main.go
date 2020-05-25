package main

import "fmt"

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println("Area:", s.area())
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func main() {
	shapes := []shape{
		triangle{10, 20},
		triangle{1, 2},
		square{4},
	}
	for _, shape := range shapes {
		printArea(shape)
	}
}
