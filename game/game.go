package game

import (
	"github.com/apiotrowski312/scrabbleBot/board"
	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/letters"
	"github.com/apiotrowski312/scrabbleBot/rack"
)

func CreateGame(dictFile, tilesFile, boardFile string, users []string, rackSize int) game {
	root, _ := gaddag.CreateGraph(dictFile)
	tb, lv, _ := letters.LoadTilesFromFile(tilesFile)
	b, _ := board.LoadBoardFromFile(boardFile)

	scrabblePlayers := []player{}
	for _, n := range users {
		scrabblePlayers = append(scrabblePlayers, player{
			name: n,
			rack: rack.CreateRack(rackSize),
		})
	}

	return game{
		board:        *b,
		dictionary:   *root,
		letterValues: *lv,
		bag:          *tb,
		players:      scrabblePlayers,
		round:        1,
	}
}

func (g game) PlaceWord(letters string, startCord [2]int, horizontal bool) (int, error) {

	isOk, err := g.players[g.round%len(g.players)].rack.AreThereLetters([]rune(letters))
	if !isOk {
		return 0, err
	}

	word, cords := g.board.MakeProperDataFormat(letters, startCord, horizontal)

	isOk, err = g.isWordPlacedCorectly(word, cords, horizontal)

	if !isOk {
		return 0, err
	}

	score := g.countScore(word, cords, horizontal)

	g.board.PlaceWord(word, cords, horizontal)

	g.players[g.round%len(g.players)].points = g.players[g.round%len(g.players)].points + score // TODO: finish this. I guess we will need refactor some stuff
	g.round++

	g.players[g.round%len(g.players)].rack.RemoveFromRack([]rune(letters))
	g.players[g.round%len(g.players)].rack.AddToRack(g.bag.DrawTiles(len(letters)))

	return score, nil
}

func (g game) isWordPlacedCorectly(word string, startCord [2]int, horizontal bool) (bool, error) {
	isOk, err := g.dictionary.IsWordValid(word[:1] + "." + word[1:])

	if !isOk {
		return false, err
	}

	isOk, err = g.board.IsWordInProperPlace(word, startCord, horizontal)
	if !isOk {
		return false, err
	}

	words, _ := g.board.CollectAllUsedWords(word, startCord, horizontal)

	for _, newWord := range words[:len(words)-1] {
		isOk, err = g.dictionary.IsWordValid(newWord)
		if !isOk {
			return false, err
		}
	}

	return isOk, err
}

func (g game) countScore(word string, startCord [2]int, horizontal bool) int {
	words, tiles := g.board.CollectAllUsedWords(word, startCord, horizontal)

	score := g.letterValues.CountPoints(words, tiles)

	return score
}
