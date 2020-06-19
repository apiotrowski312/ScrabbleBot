package grabble

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

// TODO: Create enum for wWlLs

type Grabble struct {
	Board         board.Board
	Players       []player.Player
	Bag           bag.Bag
	LettterPoints bag.LettersPoint
	CurrentRound  int
	Dict          gaddag.Node
	RackSize      int
}

func CreateGrabble(dictionary string, b [15][15]rune, nicks []string, allTiles []rune, tilePoints map[rune]int, rackSize int) Grabble {
	board := board.CreateBoard(b)
	dict, _ := gaddag.CreateGraph(dictionary)
	ba := bag.CreateBag(allTiles)
	lp := bag.CreateLettersPoint(tilePoints)

	players := []player.Player{}
	for _, p := range nicks {
		players = append(players, player.CreatePlayer(p))
	}

	return Grabble{
		Board:         *board,
		Players:       players,
		Bag:           ba,
		Dict:          *dict,
		LettterPoints: lp,
		RackSize:      rackSize,
	}
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
	if len(letters) == g.RackSize {
		points += 50
	}
	g.CurrentPlayer().AddPoints(points)
	g.CurrentPlayer().UpdateRack(letters, g.Bag.DrawLetters(len(letters)))
	g.CurrentRound++
	return nil
}

func (g Grabble) CurrentPlayer() *player.Player {
	return &g.Players[g.CurrentRound%len(g.Players)]
}

func (g *Grabble) PassTurn() {
	g.CurrentRound++
}
