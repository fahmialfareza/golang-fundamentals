package main

import "fmt"

func main() {
	// for i := 1; i < 100; i++ {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// }

	// i := 1
	// for i < 100 {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// 	i++
	// }

	// title := "Golang the best language"

	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }

	title := "Golang the best language"

	for index, letter := range title {
		if index % 2 == 0 {
			fmt.Println("letter :", string(letter))
		}
	}

	for _, letter := range title {
		letterString := string(letter)

		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("letter :", letterString)
		// }

		switch letterString {
			case "a", "i", "u", "e", "o":
				fmt.Println("letter :", letterString)
		}
	}
}