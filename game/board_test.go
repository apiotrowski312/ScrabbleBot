package game

import (
	"errors"
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

func Test_loadBoardFromFile(t *testing.T) {
	b, err := loadBoardFromFile("../exampleData/scrable.board")

	var expectedBoard board
	test_utils.BytesToStruct(t, test_utils.GetGoldenFileJSON(t, b, t.Name(), *update), &expectedBoard)

	assert.Equal(t, nil, err)
	assert.Equal(t, &expectedBoard, b, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedBoard, b))
}

func Test_isWordInProperPlace(t *testing.T) {
	var testBoard board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

	isOk, err := testBoard.isWordInProperPlace("book", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk, fmt.Sprintf("Expected board: \n%v, got: \n%v", true, isOk))

	isOk, err = testBoard.isWordInProperPlace("book", [2]int{1, 0}, true)

	assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
	assert.Equal(t, false, isOk, fmt.Sprintf("Expected board: \n%v, got: \n%v", false, isOk))
}

func Test_placeWord(t *testing.T) {
	var testBoard board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

	testBoard.placeWord("book", [2]int{2, 0}, true)

	var expectedBoard board
	test_utils.BytesToStruct(t, test_utils.GetGoldenFileJSON(t, testBoard, t.Name(), *update), &expectedBoard)

	assert.Equal(t, expectedBoard, testBoard, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedBoard, testBoard))
}

func Test_collectAllUsedWords(t *testing.T) {
	testBoard := board{
		[]tile{tile{TileType: 'W', Letter: 'g'}, tile{TileType: '0'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W', Letter: 'b'}},
		[]tile{tile{TileType: '0', Letter: 'o'}, tile{TileType: 'L'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0', Letter: 'o'}},
		[]tile{tile{TileType: 'W', Letter: 'd'}, tile{TileType: '0'}, tile{TileType: 's'}, tile{TileType: '0'}, tile{TileType: 'W', Letter: 's'}},
		[]tile{tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0', Letter: 's'}},
		[]tile{tile{TileType: 'W'}, tile{TileType: '0'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W', Letter: 'x'}},
	}

	expected := []string{"socks", "sdog."}
	expectedTiles := []string{"0L0L0", "0"}
	words, tiles := testBoard.collectAllUsedWords("socks", [2]int{3, 0}, true)
	assert.Equal(t, expected, words, fmt.Sprintf("Expected board: \n%v, got: \n%v", expected, words))
	assert.Equal(t, expectedTiles, tiles, fmt.Sprintf("Expected board: \n%v, got: \n%v", expectedTiles, tiles))
}
