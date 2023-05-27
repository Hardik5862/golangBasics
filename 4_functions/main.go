package main

import (
	"fmt"
)

func greet(name string) string {
	return "Hi! " + name
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greet("Hardik"))
	fmt.Println(sum(34, 5))
}
