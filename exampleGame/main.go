package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
)

var loopNumber = flag.Int("times", 1, "number of games to play")
var screenshot = flag.Bool("screenshot", false, "create screenshot after each round")
var winnerScreenshot = flag.Bool("winshot", false, "create screenshot with finished game")

func main() {
	game := grabble.CreateDefaultGame([][2]string{{"Bot 1", "default"}, {"Bot 2", "default"}})
	flag.Parse()
	for x := 0; x < *loopNumber; x++ {
		Game(&game)
	}
}

func Game(game *grabble.Grabble) {
	if *screenshot == true {
		img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
	}
	for !game.Stats.Finished {
		bestWord, err := game.PickBestWord()
		if err != nil {
			fmt.Println("Round passed", game.Stats.CurrentRound)
			game.PassTurn()
		}

		err = game.PlaceWord(bestWord.Word, bestWord.Cords, bestWord.Horizontal)

		if err != nil {
			fmt.Println("Round passed", game.Stats.CurrentRound)
			game.PassTurn()
		}

		if *screenshot == true {
			img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/round_%v.png", game.Stats.CurrentRound))
		}
	}

	if *winnerScreenshot == true {
		img_printer.PrintScreenBoard(*game, fmt.Sprintf("./img/finished-%v.png", time.Now().UnixNano()))
	}
	fmt.Printf("%v - Winner: %v\tPoints: %v\t Turns: %v\n", time.Now().Format("2006-01-02_15:04:05.000000"), game.Stats.Winner.Name, game.Stats.Winner.Points, game.Stats.CurrentRound)
}
