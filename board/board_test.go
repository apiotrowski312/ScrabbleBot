package board_test

import (
	"errors"
	"flag"
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/board"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_LoadBoardFromFile(t *testing.T) {
	b, err := board.LoadBoardFromFile("../exampleData/scrable.board")

	var expectedBoard board.Board
	test_utils.GetGoldenFileJSON(t, b, &expectedBoard, t.Name(), *update)

	assert.Equal(t, nil, err)
	assert.Equal(t, &expectedBoard, b)
}

func Test_IsWordInProperPlace(t *testing.T) {
	var testBoard board.Board

	t.Run("Empty board tests", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

		isOk, err := testBoard.IsWordInProperPlace("book", [2]int{2, 0}, true)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)

		isOk, err = testBoard.IsWordInProperPlace("book", [2]int{1, 0}, true)
		assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
		assert.Equal(t, false, isOk)

		isOk, err = testBoard.IsWordInProperPlace("sos", [2]int{2, 2}, false)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)

		isOk, err = testBoard.IsWordInProperPlace("sos", [2]int{2, 1}, false)
		assert.Equal(t, errors.New("There is no hooks. Wrong place"), err)
		assert.Equal(t, false, isOk)
	})

	t.Run("Chair and above board tests", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/chair_above_5x5.fixture", &testBoard)

		isOk, err := testBoard.IsWordInProperPlace("book", [2]int{2, 0}, true)
		assert.Equal(t, errors.New("You can't overwrite letter"), err)
		assert.Equal(t, false, isOk)

		isOk, err = testBoard.IsWordInProperPlace("love", [2]int{1, 2}, false)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})

	t.Run("All love board tests", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &testBoard)

		isOk, err := testBoard.IsWordInProperPlace("word", [2]int{1, 4}, false)
		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})

}

func Test_PlaceWord(t *testing.T) {
	var testBoard board.Board

	t.Run("Horizontal", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

		testBoard.PlaceWord("book", [2]int{2, 0}, true)

		var expectedBoard board.Board
		test_utils.GetGoldenFileJSON(t, testBoard, &expectedBoard, t.Name(), *update)

		assert.Equal(t, expectedBoard, testBoard)
	})

	t.Run("Vertical", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

		testBoard.PlaceWord("book", [2]int{1, 2}, false)

		var expectedBoard board.Board
		test_utils.GetGoldenFileJSON(t, testBoard, &expectedBoard, t.Name(), *update)

		assert.Equal(t, expectedBoard, testBoard)
	})
}

func Test_CollectAllUsedWords(t *testing.T) {
	var testBoard board.Board
	t.Run("Horizontal 1", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

		expected := []string{"socks"}
		expectedTiles := []string{"0L0L0"}
		words, tiles := testBoard.CollectAllUsedWords("socks", [2]int{3, 0}, true)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 1", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &testBoard)

		expected := []string{"socks"}
		expectedTiles := []string{"W0W0W"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{0, 0}, false)
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 2", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/chair_above_5x5.fixture", &testBoard)

		expected := []string{"raw"}
		expectedTiles := []string{"00w"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{4, 0}, true)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 2", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/chair_above_5x5.fixture", &testBoard)

		expected := []string{"love"}
		expectedTiles := []string{"000w"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{1, 2}, false)
		assert.Equal(t, expected, words, fmt.Sprintf("%v, %v ", expected, words))
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 3", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/hair_raw_5x5.fixture", &testBoard)

		expected := []string{"rwar.", "car"}
		expectedTiles := []string{"W", "w0W"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{4, 2}, true)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 3", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/hair_raw_5x5.fixture", &testBoard)

		expected := []string{"c.hair", "cat"}
		expectedTiles := []string{"0", "0W0"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{1, 0}, false)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Horizontal 4", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &testBoard)

		expected := []string{"c.all", "car"}
		expectedTiles := []string{"W", "W0w"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{0, 0}, true)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})

	t.Run("Vertical 4", func(t *testing.T) {
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &testBoard)

		expected := []string{"revol.", "word"}
		expectedTiles := []string{"0", "0W0W"}
		words, tiles := testBoard.CollectAllUsedWords(expected[len(expected)-1], [2]int{1, 4}, false)
		assert.Equal(t, expected, words)
		assert.Equal(t, expectedTiles, tiles)
	})
}

// testBoard = board.Board{
// 	[]board.Tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
// 	[]board.Tile{{TileType: '0'}, {TileType: 'L', Letter: 'h'}, {TileType: '0', Letter: 'a'}, {TileType: 'L', Letter: 'i'}, {TileType: '0', Letter: 'r'}},
// 	[]board.Tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 's'}, {TileType: '0'}, {TileType: 'W', Letter: 'a'}},
// 	[]board.Tile{{TileType: '0'}, {TileType: 'L'}, {TileType: '0'}, {TileType: 'L'}, {TileType: '0', Letter: 'w'}},
// 	[]board.Tile{{TileType: 'W'}, {TileType: '0'}, {TileType: 'w'}, {TileType: '0'}, {TileType: 'W'}},
// }
// test_utils.GetGoldenFileJSON(t, testBoard, t.Name(), true)
