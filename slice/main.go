package main

import "fmt"

func main() {
	// var gamingConsoles []string

	// gamingConsoles = append(gamingConsoles, "Play Station 4")
	// gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	// gamingConsoles = append(gamingConsoles, "Xbox One")

	gamingConsoles := []string{"Play Station 4", "Nintendo Switch", "Xbox One"}

	// fmt.Println(gamingConsoles)

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}