package game

import (
	"github.com/apiotrowski312/scrabbleBot/gaddag"
)

type game struct {
	bag          tileBag
	letterValues letterValue
	board        board
	dictionary   gaddag.Node // to do. Unexport struct, leave type only
}

func (g game) PlaceWord(word string, startCord [2]int, horizontal bool) (int, error) {
	isOk, err := g.dictionary.IsWordValid(word[:1] + "." + word[1:])

	if !isOk {
		return 0, err
	}

	isOk, err = g.isWordPlacedCorectly(word, startCord, horizontal)

	if !isOk {
		return 0, err
	}

	g.board.placeWord(word, startCord, horizontal)

	return 0, nil
}

func (g game) isWordPlacedCorectly(word string, startCord [2]int, horizontal bool) (bool, error) {
	isOk, err := g.board.isWordInProperPlace(word, startCord, horizontal)
	if !isOk {
		return false, err
	}

	additionalWords := g.board.collectAllUsedWords(word, startCord, horizontal)

	for _, newWord := range additionalWords {
		isOk, err = g.dictionary.IsWordValid(newWord)
		if !isOk {
			return false, err
		}
	}

	return isOk, err
}
