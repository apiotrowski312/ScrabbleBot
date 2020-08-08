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

	for !game.Stats.Finished && game.Stats.CurrentRound < 20 {
		word := game.PickBestWord()
		fmt.Println("Current player:", game.CurrentPlayer().Name)
		fmt.Println("Rack:", game.CurrentPlayer().Rack)
		fmt.Println("Points:", game.CurrentPlayer().Points)

		if word.Word == "" {
			game.PassTurn()
			fmt.Println("Turn passed")
		} else {
			err := game.PlaceWord(word.Word, word.Cords, word.Horizontal)
			fmt.Println("Word to place:", word.Word, word.Cords, word.Horizontal, word.Points)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println()

		img_printer.PrintScreenBoard(game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
	}

}
