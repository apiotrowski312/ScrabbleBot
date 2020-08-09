package grabble

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

// TODO: Create one place with all fixtures etc. to make it easier to manage
// TODO: Create enum for wWlLs
type gameStats struct {
	CurrentRound int
	Finished     bool
	Winner       *player.Player
}

// FIXME: gaddag finds words with too much lettersw

type Grabble struct {
	Board            board.Board
	Players          []player.Player
	Bag              bag.Bag
	LettterPoints    bag.LettersPoint
	Dict             gaddag.Node
	RackSize         int
	passedTurnInARow int

	Stats gameStats
}

// FIXME: You cant place word with hook like this (W in SW would be a hook):
// WORDS
//     WORDS

func CreateGrabble(dictionary string, b [15][15]rune, nicks []string, allTiles []rune, tilePoints map[rune]int, rackSize int) Grabble {
	board := board.CreateBoard(b)
	dict, _ := gaddag.CreateGraph(dictionary)
	ba := bag.CreateBag(allTiles)
	lp := bag.CreateLettersPoint(tilePoints)

	players := []player.Player{}
	for _, p := range nicks {
		players = append(players, player.CreatePlayer(p))
	}
	game := Grabble{
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

	for i := range game.Players {
		game.Players[i].UpdateRack([]rune{}, game.Bag.DrawLetters(rackSize))
	}
	log.Debug("")
	log.Debug("Grabble game created")
	return game
}

// PlaceWord - this is the most important function to use. This function will:
// - validate all created words by this move,
// - look for letter conflicts
// - increment round counter,
// - update player rack,
// - check if game should still going.
// Pass word which you want to create.
func (g *Grabble) PlaceWord(word string, startPos [2]int, horizontal bool) error {
	log.Debugf("PlaceWord function called by %s", g.CurrentPlayer().Name)

	letters, err := g.validateAndExtractUsedNewLetters(word, startPos, horizontal)
	if err != nil {
		log.Debugf("Word %v(%v) is not valid. Error: %v", word, startPos, err)
		return err
	}

	points, err := g.countPoints(word, len(letters), startPos, horizontal)
	if err != nil {
		log.Debugf("Word %v(%v) was invalid when counting points. Error: %v", word, startPos, err)
		return err
	}

	err = g.CurrentPlayer().UpdateRack(letters, g.Bag.DrawLetters(len(letters)))
	if err != nil {
		log.Debugf("Smth went wrong while updating rack. Error: %v", err)
		return err
	}

	g.Board.PlaceWord(word, startPos, horizontal)
	g.CurrentPlayer().AddPoints(points)
	g.shouldGameEnd()
	g.Stats.CurrentRound++
	g.passedTurnInARow = 0
	return nil
}

func (g Grabble) validateAndExtractUsedNewLetters(word string, startPos [2]int, horizontal bool) ([]rune, error) {
	letters, isOk := g.Board.DoesHookExist(word, startPos, horizontal)
	if isOk == false {
		return []rune{}, fmt.Errorf("word %v cannot be placed here (%v)", word, startPos)
	}

	if err := g.CurrentPlayer().AreLettersInRack(letters); err != nil {
		return []rune{}, err
	}

	if len(letters) == 0 {
		return []rune{}, fmt.Errorf("no new letters would be use with this word (%v)", word)
	}

	return letters, nil
}

func (g Grabble) countPoints(word string, numOfUsedLetters int, startPos [2]int, horizontal bool) (int, error) {
	words, bonuses := g.Board.GetAllWordsAndBonuses(word, startPos, horizontal)

	for _, word := range words {
		if isValid, err := g.Dict.IsWordValid(word); isValid == false {
			return 0, err
		}
	}

	points := g.LettterPoints.GetPoints(words, bonuses)

	// If all letters were used, add bonus 50 points (Scrabble)
	if numOfUsedLetters == g.RackSize {
		points += 50
	}
	return points, nil
}

// CurrentPlayer - return player who should do a move next.
func (g Grabble) CurrentPlayer() *player.Player {
	return &g.Players[g.Stats.CurrentRound%len(g.Players)]
}

// PassTurn - player omit his turn. No points for him.
// TODO: count number of passed turns, if its 2 for all players - finish game.
func (g *Grabble) PassTurn() {
	log.Debugf("PassTurn function called by %s. Turn %v", g.CurrentPlayer().Name, g.Stats.CurrentRound)
	g.Stats.CurrentRound++
	g.passedTurnInARow++
	g.shouldGameEnd()
}

// ChangeTiles - change tiles. Important - player will lost turn.
func (g *Grabble) ChangeTiles(tilesToChange []rune) {
	log.Debugf("ChangeTiles function called by %s", g.CurrentPlayer().Name)
	if err := g.CurrentPlayer().UpdateRack(tilesToChange, g.Bag.ChangeLetters(tilesToChange)); err != nil {
		log.Debugf("Smth went wrong while changing tiles. Error: %v", err)
		return
	}
	g.Stats.CurrentRound++
}

func (g *Grabble) shouldGameEnd() {

	// FIXME: fix below code. it doesnt work as intended now
	if g.passedTurnInARow >= len(g.Players)*2 {
		g.Stats.Finished = true
		return
	}

	// Part for ending game because lack of tiles
	if len(g.Bag) != 0 {
		return
	}

	if len(g.CurrentPlayer().Rack) != 0 {
		return
	}

	g.finishGameNoTilesLeft()
}

// FIXME: Players should have minus points if they have letters on tile.
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

	log.Debugf("Game finished. Winner is %s\n with %v points", g.Stats.Winner.Name, g.Stats.Winner.Points)
}
