package main

import (
	"fmt"
)

func main() {
	a := 4
	b := &a

	fmt.Println(a, b, *b, *&a)
	fmt.Printf("%T\n", b)
}
