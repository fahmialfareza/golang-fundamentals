package main

import (
	"errors"
	"fmt"
)

func main() {
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println(total)

	result, err := calculate(10, 2, "+")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	// result, err := calculate(10, 2, "=")
}

func calculate(number int, numberTwo int, operation string) (result int, err error) {
	switch operation {
		case "+":
			result = number + numberTwo
		case "-":
			result = number - numberTwo
		case "*":
			result = number * numberTwo
		case "/":
			result = number / numberTwo
		default:
			err = errors.New("unknown operation")
	}

	return
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}

	return total
}