package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(8, 9)

	fmt.Println(result)

	// Quiz
	resultQuiz := calculation.Multiply(4, 5)

	fmt.Println(resultQuiz)
}