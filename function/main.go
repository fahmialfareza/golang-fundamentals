package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang belajar Golang")
	fmt.Println(sentence)

	result := add(10, 20)
	fmt.Println(result)

	// luas, keliling := calculate(10, 10)
	luas, _ := calculate(10, 10)
	fmt.Println(luas)
	// fmt.Println(keliling)
}

func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}

// func calculate(panjang int, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)

// 	return luas, keliling
// }

func add(number int, numberTwo int) int {
	return number + numberTwo
}

func printMyResult(sentence string) string {
	fmt.Println(sentence)

	return sentence
}
