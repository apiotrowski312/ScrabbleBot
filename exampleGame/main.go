package main

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/grabble"
)

func main() {
	Game()
}

func Game() {

	for x := 0; x < 1000; x++ {
		game := grabble.CreateDefaultGame([]string{"Bot 1", "Bot 2"})
		for !game.Stats.Finished {
			bestWords := game.PickBestWord(50)
			wordPlaced := false
			for _, word := range bestWords {
				err := game.PlaceWord(word.Word, word.Cords, word.Horizontal)
				if err == nil {
					wordPlaced = true
					break
				}
			}

			if !wordPlaced {
				game.PassTurn()
			}

			// img_printer.PrintScreenBoard(game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
		}
		fmt.Printf("Winner: %v\tPoints: %v\t Turns: %v\n", game.Stats.Winner.Name, game.Stats.Winner.Points, game.Stats.CurrentRound)
	}

}
