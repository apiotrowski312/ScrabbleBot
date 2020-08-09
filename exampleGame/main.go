package main

import (
	"flag"
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
)

func main() {
	Game()
}

var loopNumber = flag.Int("times", 1, "number of games to play")
var screenshot = flag.Bool("screenshot", false, "should I create jpg output")

func Game() {
	flag.Parse()
	for x := 0; x < *loopNumber; x++ {
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
			if *screenshot == true {
				img_printer.PrintScreenBoard(game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
			}
		}
		fmt.Printf("Winner: %v\tPoints: %v\t Turns: %v\n", game.Stats.Winner.Name, game.Stats.Winner.Points, game.Stats.CurrentRound)
	}

}
