package main

import (
	"fmt"
)

func main() {
	// loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Arrays
	var testArr [2]string
	testArr[0] = "quiz"
	testArr[1] = "mon"
	fmt.Println(testArr)

	testArr2 := [2]string{"mcq", "tue"}
	fmt.Println(testArr2)

	testSlice := []string{"mcq", "tue", "1", "2"}
	fmt.Println(testSlice)
	fmt.Println(len(testSlice), testSlice[1:3])

	// maps
	emailList := make(map[string]string)
	emailList["Sharon"] = "sharon@gmail.com"
	emailList["Luna"] = "snow@gmail.com"
	emailList["Dark"] = "hawed@night.com"
	fmt.Println(len(emailList), emailList)

	delete(emailList, "Dark")
	fmt.Println(emailList)

	chars := map[string]string{"Sharon": "marvel", "Luna": "marvel", "Flash": "DC"}
	fmt.Println(chars)

	// range
	for i, val := range testSlice {
		fmt.Printf("%d - %s\n", i, val)
	}

	for key, val := range emailList {
		fmt.Printf("%s = %s\n", key, val)
	}
}
