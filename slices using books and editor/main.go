package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "Drakulla"
	books[1] = "Pyasa Darinda"
	books[2] = "Island"

	newBooks := [5]string{"uselessness", "Strong"}

	books = newBooks

	games := []string{"kokemon", "sims"}

	newGames := []string{"GTA5", "Shrimops"}

	newGames = games

	games = nil

	var ok string

	for i, game := range games {
		if game != newGames[i] {
			ok = "not"
			break

		}
	}

	if len(games) == len(newGames) {
		ok = "not"
	}

	fmt.Printf("games and newGames are %sequal\n\n", ok)

	fmt.Printf("books        : %#v\n", books)
	fmt.Printf("newBooks     : %#v\n", newBooks)
	fmt.Printf("Games        : %#v\n", games)
	fmt.Printf("games        : %T\n", games)
	fmt.Printf("Games Lenght : %d\n", len(games))
	fmt.Printf("Games equal nil : %t\n", games == nil)

}
