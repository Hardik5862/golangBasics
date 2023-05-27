package main

import (
	"fmt"
)

type Character struct {
	name     string
	universe string
	gender   string
	rating   int
}

// method - value receiver
func (c Character) intro() string {
	return "Hi! I'm " + c.name + " from " + c.universe + " universe."
}

// method - pointer receiver for changing object
func (c *Character) increaseRating() {
	c.rating++
}

func main() {
	char1 := Character{
		name:     "Sharon Rogers",
		universe: "marvel",
		gender:   "F",
		rating:   10,
	}

	fmt.Println(char1, char1.name) // can mutate toooo
	fmt.Println(char1.intro())

	char1.increaseRating()
	fmt.Println("Rating (/10): ", char1.rating)
}
