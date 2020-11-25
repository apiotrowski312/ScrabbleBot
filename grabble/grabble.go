package grabble

import (
	"fmt"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

// TODO: Create enum for wWlLs
type gameStats struct {
	CurrentRound int
	Finished     bool
	Winner       *player.Player
}

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

// Simple singleton. Just for testing purposes as it would take a lot of time to generate graph
// e.g. 1000 times
var gaddagFullGraph *gaddag.Node

func CreateGrabble(dictionary string, b [15][15]rune, nicks [][2]string, allTiles []rune, tilePoints map[rune]int, rackSize int) Grabble {
	board := board.CreateBoard(b)
	if gaddagFullGraph == nil {
		gaddagFullGraph, _ = gaddag.CreateGraph(dictionary)
	}
	ba := bag.CreateBag(allTiles)
	lp := bag.CreateLettersPoint(tilePoints)

	players := []player.Player{}
	for _, p := range nicks {
		players = append(players, player.CreatePlayer(p[0], p[1]))
	}
	game := Grabble{
		Board:         *board,
		Players:       players,
		Bag:           ba,
		Dict:          *gaddagFullGraph,
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
	log.Tracef("PlaceWord function called by %s", g.CurrentPlayer().Name)

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

	g.Board.PlaceWord(word, startPos, g.CurrentPlayer().Name, g.Stats.CurrentRound, horizontal)
	g.CurrentPlayer().AddPoints(points)
	log.Infof("Player %v placed word %v(cords:%v, horizontal: %v) on the board. Sum value of %v points", g.CurrentPlayer(), word, startPos, horizontal, points)
	g.shouldGameEnd()
	g.nextRound(false)
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

	// If all letters were used, add 50 points bonus
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
func (g *Grabble) PassTurn() {
	log.Tracef("PassTurn function called by %s. Turn %v", g.CurrentPlayer().Name, g.Stats.CurrentRound)
	g.nextRound(true)
	g.shouldGameEnd()
}

// ChangeTiles - change tiles. Important - player will lost turn.
func (g *Grabble) ChangeTiles(tilesToChange []rune) {
	log.Tracef("ChangeTiles function called by %s", g.CurrentPlayer().Name)
	if len(g.Bag) == 0 {
		// FIXME: Return error with information about no tiles in bag.
		log.Debugf("No tiles in bag. Failover to PassTurn")
		g.PassTurn()
		return
	}
	if err := g.CurrentPlayer().UpdateRack(tilesToChange, g.Bag.ChangeLetters(tilesToChange)); err != nil {
		log.Debugf("Smth went wrong while changing tiles. Error: %v", err)
		return
	}
	g.nextRound(false)
}

func (g *Grabble) shouldGameEnd() {

	if g.passedTurnInARow >= len(g.Players)*2 {
		g.finishGame()
		return
	}

	if len(g.Bag) == 0 && len(g.CurrentPlayer().Rack) == 0 {
		g.finishGame()
		return
	}

}

// FIXME: Change naming of this function
func (g *Grabble) nextRound(roundPassed bool) {
	if roundPassed {
		g.passedTurnInARow++
	} else {
		g.passedTurnInARow = 0
	}

	g.Stats.CurrentRound++
	log.Debugf("Current round: %v", g.Stats.CurrentRound)
}

func (g *Grabble) finishGame() {
	pointsFromRacks := 0

	for _, p := range g.Players {
		bonus := ""
		for range p.Rack {
			bonus += "0"
		}
		points := g.LettterPoints.GetPoints([]string{string(p.Rack)}, []string{bonus})
		pointsFromRacks += points
		p.MinusPoints(points)
	}

	if len(g.CurrentPlayer().Rack) == 0 {
		g.CurrentPlayer().AddPoints(pointsFromRacks)
	}

	g.Stats.Finished = true
	for i, p := range g.Players {
		if p.Points > g.Stats.Winner.Points {
			log.Debugf("Player %v(%v) has more points than %v(%v)", p.Name, p.Points, g.Stats.Winner.Name, g.Stats.Winner.Points)
			g.Stats.Winner = &g.Players[i]
			log.Debugf("Now %v is winning", g.Stats.Winner.Name)
		}
	}

	log.Infof("Game finished. Winner is %s with %v points", g.Stats.Winner.Name, g.Stats.Winner.Points)
}
