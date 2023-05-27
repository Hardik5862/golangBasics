package main

import (
	"fmt"
)

func main() {
	// if/else
	x, y := 10, 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Println("Both are equal")
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// switch
	n := 1

	switch n % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Never")
	}
}
