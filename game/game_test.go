package game

import (
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/bmizerany/assert"
)

func Test_PlaceWord(t *testing.T) {
	root, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	lV := letterValue{
		'_': 0, 'e': 1, 'a': 1, 'i': 1, 'o': 1, 'n': 1,
		'r': 1, 't': 1, 'l': 1, 's': 1, 'u': 1, 'd': 2,
		'g': 2, 'b': 3, 'c': 3, 'm': 3, 'p': 3, 'f': 4,
		'h': 4, 'v': 4, 'w': 4, 'y': 4, 'k': 5, 'j': 8,
		'x': 8,
	}

	gameForTest := game{
		board: board{
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		},
		dictionary:   *root,
		letterValues: lV,
	}

	gameAfterPlaceWord := game{
		board: board{
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W', letter: 'b'}, tile{tileType: '0', letter: 'o'}, tile{tileType: 's', letter: 'o'}, tile{tileType: '0', letter: 'k'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		},
		dictionary:   *root,
		letterValues: lV,
	}

	score, err := gameForTest.PlaceWord("book", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, gameAfterPlaceWord, gameForTest, fmt.Sprintf("Expected board: \n%v, got: \n%v", gameAfterPlaceWord, gameForTest))
	assert.Equal(t, 30, score, fmt.Sprintf("Expected score: \n%v, got: \n%v", 29, score))
}

func Test_Game_isWordPlacedCorectly(t *testing.T) {

	root, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	gameForTest := game{
		board: board{
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		},
		dictionary: *root,
	}

	isOk, err := gameForTest.isWordPlacedCorectly("book", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk, fmt.Sprintf("Expected score: \n%v, got: \n%v", true, isOk))

	gameForTest = game{
		board: board{
			[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W', letter: 's'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L', letter: 'b'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0', letter: 'o'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0', letter: 'o'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
			[]tile{tile{tileType: '0'}, tile{tileType: 'L', letter: 's'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
			[]tile{tile{tileType: 'W'}, tile{tileType: '0', letter: 's'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		},
		dictionary: *root,
	}

	isOk, err = gameForTest.isWordPlacedCorectly("books", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk, fmt.Sprintf("Expected score: \n%v, got: \n%v", true, isOk))
}
