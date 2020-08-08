package main

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
)

func main() {
	Game()
}

func Game() {
	game := grabble.CreateDefaultGame([]string{"Zuza", "Olek"})

	for !game.Stats.Finished {
		bestWords := game.PickBestWord(50)
		fmt.Println("Current player:", game.CurrentPlayer().Name)
		fmt.Println("Rack:", game.CurrentPlayer().Rack)
		fmt.Println("Points:", game.CurrentPlayer().Points)

		wordPlaced := false
		for _, word := range bestWords {
			err := game.PlaceWord(word.Word, word.Cords, word.Horizontal)
			if err == nil {
				wordPlaced = true
				fmt.Println("Word to place:", word.Word, word.Cords, word.Horizontal, word.Points)
				break
			} else {
				fmt.Println(err)
			}
		}

		if !wordPlaced {
			game.PassTurn()
			fmt.Println("Pass turn")
		}
		fmt.Println()

		img_printer.PrintScreenBoard(game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
	}

	fmt.Printf("Winner is %v, with %v points\nIn %v rounds\n", game.Stats.Winner.Name, game.Stats.Winner.Points, game.Stats.CurrentRound)
}
