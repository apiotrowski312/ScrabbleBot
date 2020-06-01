package grabble_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_CreateBoard(t *testing.T) {
	type testCase struct {
		name     string
		template [15][15]rune
	}
	test := []testCase{
		{
			"Proper template",
			[15][15]rune{
				{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'},
				{'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0'},
				{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'},
			},
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedBoard grabble.Board
			board := grabble.CreateBoard(c.template)
			test_utils.GetGoldenFileJSON(t, board, &expectedBoard, c.name, *update)

			assert.Equal(t, &expectedBoard, board)
		})
	}
}

func Test_TransposeBoard(t *testing.T) {
	var board grabble.Board
	test_utils.LoadJSONFixture(t, "testdata/board.fixture", &board)

	transposedBoard := board.TransposeBoard()
	board[5][7].Letter = 'c'
	assert.Equal(t, board[5][7], transposedBoard[7][5])

	transposedBoard[5][7].Letter = 'a'
	assert.Equal(t, transposedBoard[5][7], board[7][5])

	var expectedTransposedBoard grabble.Board
	test_utils.GetGoldenFileJSON(t, transposedBoard, &expectedTransposedBoard, "transposed_board", *update)
	assert.Equal(t, &expectedTransposedBoard, transposedBoard)

}

func Test_PlaceWord(t *testing.T) {
	type testCase struct {
		name       string
		words      []string
		startPos   [][2]int
		horizontal bool
	}
	test := []testCase{
		{
			"Place one word",
			[]string{"words"},
			[][2]int{
				{0, 0},
			},
			true,
		},
		{
			"Place many word",
			[]string{
				"words", "words", "words", "words", "words", "words",
			},
			[][2]int{
				{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
			},
			true,
		},
		{
			"Simple vertically",
			[]string{
				"words",
			},
			[][2]int{
				{0, 0},
			},
			false,
		},
		{
			"Words vertically",
			[]string{
				"words",
				"words",
				"words",
			},
			[][2]int{
				{0, 0}, {0, 1}, {2, 2},
			},
			false,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var board grabble.Board
			test_utils.LoadJSONFixture(t, "testdata/board.fixture", &board)
			var expectedBoard grabble.Board

			for i, word := range c.words {
				board.PlaceWord(word, c.startPos[i], c.horizontal)
			}

			test_utils.GetGoldenFileJSON(t, board, &expectedBoard, c.name, *update)
			assert.Equal(t, expectedBoard, board)
		})
	}
}
