package main

import "fmt"

func main() {
	inputs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i <= 10; i++ {
		if inputs[i]%2 == 0 {
			fmt.Println(inputs[i], "is even")
		} else {
			fmt.Println(inputs[i], "is odd")
		}
	}
	// 这样写也可以，两种写法都是对的
	// for _, number := range inputs {
	// 	if number%2 == 0 {
	// 		fmt.Println(number, "is even")
	// 	} else {
	// 		fmt.Println(number, "is odd")
	// 	}
	// }
}
