package game

import (
	"errors"
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/board"
	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/letters"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func createGameTestHelper(dictFile, tilesFile, boardFile string) game {
	root, _ := gaddag.CreateGraph(dictFile)
	tb, lv, _ := letters.LoadTilesFromFile(tilesFile)
	b, _ := board.LoadBoardFromFile(boardFile)

	return game{
		board:        *b,
		dictionary:   *root,
		letterValues: *lv,
		bag:          *tb,
	}
}

func Test_PlaceWord(t *testing.T) {
	t.Run("Place simple word", func(t *testing.T) {
		gameForTest := createGameTestHelper("../exampleData/tiny_english.txt", "../exampleData/english_tiles.csv", "../exampleData/scrable.board")

		score, err := gameForTest.PlaceWord("book", [2]int{2, 0}, true)

		var expectedGame game
		test_utils.GetGoldenFileJSON(t, gameForTest, &expectedGame, t.Name(), *update)

		assert.Equal(t, nil, err)
		assert.Equal(t, expectedGame.bag, gameForTest.bag)
		assert.Equal(t, expectedGame.board, gameForTest.board)
		assert.Equal(t, expectedGame.letterValues, gameForTest.letterValues)
		assert.Equal(t, 30, score)

	})

	t.Run("Test correctly placed word", func(t *testing.T) {
		gameForTest := createGameTestHelper("../exampleData/tiny_english.txt", "../exampleData/english_tiles.csv", "../exampleData/scrable.board")
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &gameForTest.board)

		score, err := gameForTest.PlaceWord("word", [2]int{1, 4}, false)

		var expectedGame game
		test_utils.GetGoldenFileJSON(t, gameForTest, &expectedGame, t.Name(), *update)

		assert.Equal(t, nil, err)
		assert.Equal(t, expectedGame.bag, gameForTest.bag)
		assert.Equal(t, expectedGame.board, gameForTest.board)
		assert.Equal(t, expectedGame.letterValues, gameForTest.letterValues)
		assert.Equal(t, 80, score)
	})
}

func Test_Game_IsWordPlacedCorectly(t *testing.T) {

	t.Run("Test correctly placed word", func(t *testing.T) {
		gameForTest := createGameTestHelper("../exampleData/tiny_english.txt", "../exampleData/english_tiles.csv", "../exampleData/scrable.board")
		isOk, err := gameForTest.IsWordPlacedCorectly("book", [2]int{2, 0}, true)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})

	t.Run("Test incorectly", func(t *testing.T) {
		gameForTest := createGameTestHelper("../exampleData/tiny_english.txt", "../exampleData/english_tiles.csv", "../exampleData/scrable.board")
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &gameForTest.board)

		isOk, err := gameForTest.IsWordPlacedCorectly("books", [2]int{2, 0}, true)

		assert.Equal(t, errors.New("You can't overwrite letter"), err)
		assert.Equal(t, false, isOk)
	})

	t.Run("Test correctly placed word", func(t *testing.T) {
		gameForTest := createGameTestHelper("../exampleData/tiny_english.txt", "../exampleData/english_tiles.csv", "../exampleData/scrable.board")
		test_utils.LoadJSONFixture(t, "testdata/all_love_5x5.fixture", &gameForTest.board)

		isOk, err := gameForTest.IsWordPlacedCorectly("word", [2]int{1, 4}, false)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})
}
