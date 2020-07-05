package grabble

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

// TODO: Add logging to file. It will be helpfull in future for collecting statistics.

// TODO: Create enum for wWlLs
type gameStats struct {
	CurrentRound int
	Finished     bool
	Winner       *player.Player
}

type Grabble struct {
	Board         board.Board
	Players       []player.Player
	Bag           bag.Bag
	LettterPoints bag.LettersPoint
	Dict          gaddag.Node
	RackSize      int

	Stats gameStats
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
		Stats: gameStats{
			CurrentRound: 0,
			Winner:       &player.Player{},
		},
		RackSize: rackSize,
	}
}

// PlaceWord - this is the most important function to use. This function will:
// - validate all created words by this move,
// - look for letter conflicts
// - increment round counter,
// - update player rack,
// - check if game should still going.
// Pass word which you want to create and letters you will use to do soo.
func (g *Grabble) PlaceWord(word string, letters []rune, startPos [2]int, horizontal bool) error {

	if g.Stats.Finished {
		return fmt.Errorf("game finished already")
	}

	// MAYBE: Add checker if word match letters
	if len(letters) == 0 {
		return fmt.Errorf("no letters, nothing to place")
	}

	points, err := g.countPoints(word, letters, startPos, horizontal)
	if err != nil {
		return err
	}

	g.Board.PlaceWord(word, startPos, horizontal)
	g.CurrentPlayer().AddPoints(points)
	g.CurrentPlayer().UpdateRack(letters, g.Bag.DrawLetters(len(letters)))
	g.shouldGameEnd()
	g.Stats.CurrentRound++

	return nil
}

func (g Grabble) countPoints(word string, letters []rune, startPos [2]int, horizontal bool) (int, error) {
	// FIXME: This shouldn't be here, am I right? I should change function name or move it smwhere else

	// THIS SHOULD RETURN LETTERS PLAYER NEED TO PLACE. IT WILL BE VALIDATION
	// AND THEN NO LETTERS variable is needed. It will make API easier + it will fix issiue with points :D
	if isOk := g.Board.CanWordBePlaced(word, startPos, horizontal); isOk == false {
		return 0, fmt.Errorf("word cannot be placed here")
	}

	words, bonuses := g.Board.GetAllWordsAndBonuses(word, startPos, horizontal)

	for _, word := range words {
		if isValid, err := g.Dict.IsWordValid(word); isValid == false {
			return 0, err
		}
	}

	points := g.LettterPoints.GetPoints(words, bonuses)

	// If all letters were used, add bonus 50 points (Scrabble)
	if len(letters) == g.RackSize {
		points += 50
	}
	return points, nil
}

// CurrentPlayer - return player who should do a move next.
func (g Grabble) CurrentPlayer() *player.Player {
	return &g.Players[g.Stats.CurrentRound%len(g.Players)]
}

// GetTurn - return number of current turn.
func (g Grabble) GetTurn() int {
	return g.Stats.CurrentRound%len(g.Players) + 1
}

// TODO: count number of passed turns, if its 2 for all players - finish game.
// PassTurn - player omit his turn. No points for him.
func (g *Grabble) PassTurn() {
	g.Stats.CurrentRound++
}

// ChangeTiles - change tiles. Important - player will lost turn.
func (g *Grabble) ChangeTiles(tilesToChange []rune) {
	g.CurrentPlayer().UpdateRack(tilesToChange, g.Bag.ChangeLetters(tilesToChange))
	g.Stats.CurrentRound++
}

func (g *Grabble) shouldGameEnd() {
	// Part for ending game because lack of tiles
	if len(g.Bag) != 0 {
		return
	}

	if len(g.CurrentPlayer().Rack) != 0 {
		return
	}

	g.finishGameNoTilesLeft()

	// TODO: Add ending game bacause everyboty pass round two times in a row
}

func (g *Grabble) finishGameNoTilesLeft() {
	letters := []string{}
	placeholder := []string{}

	for _, p := range g.Players {
		letters = append(letters, string(p.Rack))
	}

	for _, l := range letters {
		bonus := ""
		for range l {
			bonus += "0"
		}
		placeholder = append(placeholder, bonus)
	}

	leftTilesPoints := g.LettterPoints.GetPoints(letters, placeholder)
	g.CurrentPlayer().AddPoints(leftTilesPoints)
	g.Stats.Finished = true
	for _, p := range g.Players {
		if p.Points > g.Stats.Winner.Points {
			g.Stats.Winner = &p
		}
	}
}
