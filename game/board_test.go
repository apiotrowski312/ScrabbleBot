package game

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func Test_loadBoardFromFile(t *testing.T) {
	b, err := loadBoardFromFile("../exampleData/scrable.board")

	expectedBoard := &board{
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
	}

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedBoard, b, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedBoard, b))
}

func Test_isWordInProperPlace(t *testing.T) {
	testBoard := board{
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
	}

	isOk, err := testBoard.isWordInProperPlace("book", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk, fmt.Sprintf("Expected board: \n%v, got: \n%v", true, isOk))

	isOk, err = testBoard.isWordInProperPlace("book", [2]int{1, 0}, true)

	assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
	assert.Equal(t, false, isOk, fmt.Sprintf("Expected board: \n%v, got: \n%v", false, isOk))
}

func Test_placeWord(t *testing.T) {

	testBoard := board{
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
	}

	expectedBoard := board{
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W', letter: 'b'}, tile{tileType: '0', letter: 'o'}, tile{tileType: 's', letter: 'o'}, tile{tileType: '0', letter: 'k'}, tile{tileType: 'W'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W'}},
	}

	testBoard.placeWord("book", [2]int{2, 0}, true)
	assert.Equal(t, expectedBoard, testBoard, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedBoard, testBoard))
}

func Test_collectAllUsedWords(t *testing.T) {

	testBoard := board{
		[]tile{tile{tileType: 'W', letter: 'g'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W', letter: 'b'}},
		[]tile{tile{tileType: '0', letter: 'o'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0', letter: 'o'}},
		[]tile{tile{tileType: 'W', letter: 'd'}, tile{tileType: '0'}, tile{tileType: 's'}, tile{tileType: '0'}, tile{tileType: 'W', letter: 's'}},
		[]tile{tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0'}, tile{tileType: 'L'}, tile{tileType: '0', letter: 's'}},
		[]tile{tile{tileType: 'W'}, tile{tileType: '0'}, tile{tileType: 'w'}, tile{tileType: '0'}, tile{tileType: 'W', letter: 'x'}},
	}

	expected := []string{"socks", "sdog."}
	expectedTiles := []string{"0L0L0", "0"}
	words, tiles := testBoard.collectAllUsedWords("socks", [2]int{3, 0}, true)
	assert.Equal(t, expected, words, fmt.Sprintf("Expected board: \n%v, got: \n%v", expected, words))
	assert.Equal(t, expectedTiles, tiles, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedTiles, tiles))
}
