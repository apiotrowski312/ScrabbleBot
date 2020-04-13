package game

import (
	"encoding/json"

	"github.com/apiotrowski312/scrabbleBot/board"
	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/letters"
)

type game struct {
	bag          letters.TileBag
	letterValues letters.LetterValue
	board        board.Board
	dictionary   gaddag.Node
}

func (g game) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Bag          letters.TileBag `json:"bag"`
		LetterValues letters.LetterValue
		Board        board.Board
		Dictionary   gaddag.Node
	}{
		Bag:          g.bag,
		LetterValues: g.letterValues,
		Board:        g.board,
		Dictionary:   g.dictionary,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (g *game) UnmarshalJSON(jsonBytes []byte) error {
	type Game struct {
		Bag          letters.TileBag
		LetterValues letters.LetterValue
		Board        board.Board
		Dictionary   gaddag.Node
	}

	var exportedGame Game
	if err := json.Unmarshal(jsonBytes, &exportedGame); err != nil {
		return err
	}

	g.bag = exportedGame.Bag
	g.letterValues = exportedGame.LetterValues
	g.board = exportedGame.Board
	g.dictionary = exportedGame.Dictionary

	return nil
}

func (g game) PlaceWord(word string, startCord [2]int, horizontal bool) (int, error) {
	isOk, err := g.IsWordPlacedCorectly(word, startCord, horizontal)

	if !isOk {
		return 0, err
	}

	score := g.countScore(word, startCord, horizontal)

	g.board.PlaceWord(word, startCord, horizontal)

	return score, nil
}

func (g game) IsWordPlacedCorectly(word string, startCord [2]int, horizontal bool) (bool, error) {
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
