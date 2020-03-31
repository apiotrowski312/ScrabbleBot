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
	isOk, err := g.isWordPlacedCorectly(word, startCord, horizontal)

	if !isOk {
		return 0, err
	}

	g.board.placeWord(word, startCord, horizontal)

	score := g.countScore(word, startCord, horizontal)

	return score, nil
}

func (g game) isWordPlacedCorectly(word string, startCord [2]int, horizontal bool) (bool, error) {
	isOk, err := g.dictionary.IsWordValid(word[:1] + "." + word[1:])

	if !isOk {
		return false, err
	}

	isOk, err = g.board.isWordInProperPlace(word, startCord, horizontal)
	if !isOk {
		return false, err
	}

	words, _ := g.board.collectAllUsedWords(word, startCord, horizontal)

	for _, newWord := range words[1:] {
		isOk, err = g.dictionary.IsWordValid(newWord)
		if !isOk {
			return false, err
		}
	}

	return isOk, err
}

func (g game) countScore(word string, startCord [2]int, horizontal bool) int {
	words, tiles := g.board.collectAllUsedWords(word, startCord, horizontal)

	score := g.letterValues.countPoints(words, tiles)

	return score
}
