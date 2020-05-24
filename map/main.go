package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	var colors1 map[string]string
	fmt.Println("colors1: ", colors1)

	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"
	fmt.Println("colors2: ", colors2)

	// iterate over maps
	printMap(colors) // we are not copying the map, we are copying thre reference to the map
}

// map里面没有指针这些！

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
