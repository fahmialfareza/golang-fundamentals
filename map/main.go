package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["go"] = 9
	// myMap["javascript"] = 8
	// myMap["go"] = 10

	// myMap := map[string]string{"ruby": "is beautiful", "go": "is super fast"}

	myMap := map[string]string{
		"ruby": "is beautiful",
		"go": "is super fast",
		"javascript": "hype",
	}

	// fmt.Println(myMap["go"])

	for key, value := range myMap {
		fmt.Println("Key :", key, " Value :", value)
	}

	fmt.Println("===========")

	delete(myMap, "ruby")

	for key, value := range myMap {
		fmt.Println("Key :", key, " Value :", value)
	}

	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}