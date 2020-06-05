package grabble

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

type Grabble struct {
	Board         board.Board
	Players       []player.Player
	Bag           bag.Bag
	LettterPoints bag.LettersPoint
	CurrentRound  int
	Dict          gaddag.Node
}

func (g *Grabble) PlaceWord(word string, letters []rune, startPos [2]int, horizontal bool) error {
	if len(letters) == 0 {
		return fmt.Errorf("no letters, nothing to place")

	}

	if isOk := g.Board.CanWordBePlaced(word, startPos, horizontal); isOk == false {
		return fmt.Errorf("word cannot be placed here")
	}

	words, bonuses := g.Board.GetAllWordsAndBonuses(word, startPos, horizontal)

	for _, word := range words {
		if isValid, err := g.Dict.IsWordValid(word); isValid == false {
			return err
		}
	}

	g.Board.PlaceWord(word, startPos, horizontal)

	points := g.LettterPoints.GetPoints(words, bonuses)

	// If all letters were used, add bonus 50 points (Scrabble)
	if len(letters) == 7 {
		points += 50
	}
	g.Players[g.CurrentRound%len(g.Players)].AddPoints(points)
	g.Players[g.CurrentRound%len(g.Players)].UpdateRack(letters, g.Bag.DrawLetters(len(letters)))
	g.CurrentRound++
	return nil
}
