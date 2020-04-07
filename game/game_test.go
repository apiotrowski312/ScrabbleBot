package game

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_PlaceWord(t *testing.T) {
	root, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")
	var lv letterValue
	test_utils.LoadJSONFixture(t, "testdata/letters_values.fixture", &lv)
	var b board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &b)

	gameForTest := game{
		board:        b,
		dictionary:   *root,
		letterValues: lv,
	}

	gameAfterPlaceWord := game{
		board: board{
			[]tile{tile{TileType: 'W'}, tile{TileType: '0'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W'}},
			[]tile{tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}},
			[]tile{tile{TileType: 'W', Letter: 'b'}, tile{TileType: '0', Letter: 'o'}, tile{TileType: 's', Letter: 'o'}, tile{TileType: '0', Letter: 'k'}, tile{TileType: 'W'}},
			[]tile{tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}},
			[]tile{tile{TileType: 'W'}, tile{TileType: '0'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W'}},
		},
		dictionary:   *root,
		letterValues: lv,
	}

	score, err := gameForTest.PlaceWord("book", [2]int{2, 0}, true)

	assert.Equal(t, nil, err)
	assert.Equal(t, gameAfterPlaceWord, gameForTest)
	assert.Equal(t, 30, score)
}

func Test_Game_isWordPlacedCorectly(t *testing.T) {
	root, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")
	var lv letterValue
	test_utils.LoadJSONFixture(t, "testdata/letters_values.fixture", &lv)
	var b board
	test_utils.LoadJSONFixture(t, "testdata/empty_board_5x5.fixture", &b)

	t.Run("Test correctly placed word", func(t *testing.T) {
		gameForTest := game{
			board:        b,
			dictionary:   *root,
			letterValues: lv,
		}

		isOk, err := gameForTest.isWordPlacedCorectly("book", [2]int{2, 0}, true)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})

	t.Run("Test correctly placed word", func(t *testing.T) {
		gameForTest := game{
			board: board{
				[]tile{tile{TileType: 'W'}, tile{TileType: '0'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W', Letter: 's'}},
				[]tile{tile{TileType: '0'}, tile{TileType: 'L', Letter: 'b'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0', Letter: 'o'}},
				[]tile{tile{TileType: 'W'}, tile{TileType: '0', Letter: 'o'}, tile{TileType: 's'}, tile{TileType: '0'}, tile{TileType: 'W'}},
				[]tile{tile{TileType: '0'}, tile{TileType: 'L', Letter: 's'}, tile{TileType: '0'}, tile{TileType: 'L'}, tile{TileType: '0'}},
				[]tile{tile{TileType: 'W'}, tile{TileType: '0', Letter: 's'}, tile{TileType: 'w'}, tile{TileType: '0'}, tile{TileType: 'W'}},
			},
			dictionary: *root,
		}

		isOk, err := gameForTest.isWordPlacedCorectly("books", [2]int{2, 0}, true)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, isOk)
	})
}
