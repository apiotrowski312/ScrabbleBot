package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
)

var loopNumber = *flag.Int("times", 1, "number of games to play")
var screenshot = *flag.Bool("screenshot", false, "create screenshot after each round")
var winnerScreenshot = *flag.Bool("winshot", false, "create screenshot with finished game")

func main() {
	game := grabble.CreateDefaultGame([]string{"Bot 1", "Bot 2"})
	flag.Parse()
	for x := 0; x < loopNumber; x++ {
		Game(&game)
	}
}

func Game(game *grabble.Grabble) {
	if screenshot == true {
		img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
	}
	for !game.Stats.Finished {
		bestWords := game.PickBestWord(100)
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
		if screenshot == true {
			img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
		}
	}

	if winnerScreenshot == true {
		img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/finished-%v.png", time.Now().UnixNano()))
		fmt.Printf("%v - Winner: %v\tPoints: %v\t Turns: %v\n", time.Now().Format("2006-01-02_15:04:05.000000"), game.Stats.Winner.Name, game.Stats.Winner.Points, game.Stats.CurrentRound)
	}

}
