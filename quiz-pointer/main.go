package main

import "fmt"

type Gamer struct {
	Name	string
	Games	[]string
}

// Add Game
func (gamer *Gamer) addGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Reza"}

	gamer.addGame("Mario")
	gamer.addGame("FIFA 2022")
	gamer.addGame("Soccer 2022")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}