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
	assert.Equal(t, &expectedBoard, b)
}

func Test_isWordInProperPlace(t *testing.T) {
	var testBoard board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

	isOk, err := testBoard.isWordInProperPlace("book", [2]int{2, 0}, true)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk)

	isOk, err = testBoard.isWordInProperPlace("book", [2]int{1, 0}, true)
	assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
	assert.Equal(t, false, isOk)

	isOk, err = testBoard.isWordInProperPlace("sos", [2]int{2, 2}, false)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, isOk)

	isOk, err = testBoard.isWordInProperPlace("sos", [2]int{2, 1}, false)
	assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
	assert.Equal(t, false, isOk)
}

func Test_placeWord(t *testing.T) {
	var testBoard board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

	testBoard.placeWord("book", [2]int{2, 0}, true)

	var expectedBoard board
	test_utils.BytesToStruct(t, test_utils.GetGoldenFileJSON(t, testBoard, t.Name(), *update), &expectedBoard)

	assert.Equal(t, expectedBoard, testBoard)
}

func Test_collectOtherWordsAndTilesHorizontal(t *testing.T) {
	t.Run("Horizontal 1", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'b'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 'o'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W', Letter: 's'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 's'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'x'}},
		}

		expected := []string{"sdog."}
		expectedTiles := []string{"0"}
		words, tiles := testBoard.collectOtherWordsAndTilesHorizontal("socks", [2]int{3, 0})
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 2", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{}
		expectedTiles := []string{}
		words, tiles := testBoard.collectOtherWordsAndTilesHorizontal("socks", [2]int{3, 0})
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 3", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 's'}, {TileType: '0', Letter: 'o'}, {TileType: 'w', Letter: 'c'}, {TileType: '0', Letter: 'k'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 'a'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W', Letter: 'l'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 't'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{}
		expectedTiles := []string{}
		words, tiles := testBoard.collectOtherWordsAndTilesHorizontal("socks", [2]int{0, 0})
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})
}

func Test_collectOtherWordsAndTilesVertical(t *testing.T) {
	t.Run("Vertical 1", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{}
		expectedTiles := []string{}
		words, tiles := testBoard.collectOtherWordsAndTilesVertical("gods", [2]int{0, 0})
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 2", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0'}, {TileType: 'L', Letter: 'o'}, {TileType: '0', Letter: 's'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{"s.os"}
		expectedTiles := []string{"0"}
		words, tiles := testBoard.collectOtherWordsAndTilesVertical("gods", [2]int{0, 0})
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 3", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 's'}, {TileType: 'L', Letter: 'o'}, {TileType: '0', Letter: 's'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{}
		expectedTiles := []string{}
		words, tiles := testBoard.collectOtherWordsAndTilesVertical("test", [2]int{1, 2})
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})
}

func Test_collectAllUsedWords(t *testing.T) {
	t.Run("Horizontal 1", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'b'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 'o'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W', Letter: 's'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 's'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'x'}},
		}

		expected := []string{"sdog.", "socks"}
		expectedTiles := []string{"0", "0L0L0"}
		words, tiles := testBoard.collectAllUsedWords("socks", [2]int{3, 0}, true)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 1", func(t *testing.T) {
		testBoard := board{
			[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
			[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
			[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		}

		expected := []string{"gods"}
		expectedTiles := []string{"0000"}
		words, tiles := testBoard.collectAllUsedWords(expected[0], [2]int{0, 0}, false)
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})
}

func Test_tileUnderLayedWord(t *testing.T) {
	testBoard := board{
		[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
		[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
		[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 's'}},
		[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'x'}},
	}

	t.Run("Horizontal 1", func(t *testing.T) {
		tiles := testBoard.tileUnderLayedWord("gods", [2]int{0, 0}, false)
		expectedTiles := "0000"
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 1", func(t *testing.T) {
		tiles := testBoard.tileUnderLayedWord("door", [2]int{2, 0}, true)
		expectedTiles := "00s0"
		assert.Equal(t, expectedTiles, tiles)
	})
}

func Test_getTileType(t *testing.T) {
	testBoard := board{
		[]tile{{TileType: 'W', Letter: 'g'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
		[]tile{{TileType: '0', Letter: 'o'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0'}},
		[]tile{{TileType: 'W', Letter: 'd'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W'}},
		[]tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 's'}},
		[]tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W', Letter: 'x'}},
	}

	assert.Equal(t, '0', testBoard.getTileType([2]int{0, 0}))
	assert.Equal(t, '0', testBoard.getTileType([2]int{1, 0}))
	assert.Equal(t, '0', testBoard.getTileType([2]int{2, 0}))
	assert.Equal(t, '0', testBoard.getTileType([2]int{3, 0}))
	assert.Equal(t, 'W', testBoard.getTileType([2]int{4, 0}))

}
