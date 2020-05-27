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
		err      bool
	}
	test := []testCase{
		{
			"Proper template",
			[15][15]rune{
				[15]rune{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'},
				[15]rune{'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0'},
				[15]rune{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'},
			}, false,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedBoard grabble.Board
			board, err := grabble.CreateBoard(c.template)
			test_utils.GetGoldenFileJSON(t, board, &expectedBoard, c.name, *update)

			assert.Equal(t, expectedBoard, board)

			if c.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func Test_TransposeBoard(t *testing.T) {
	var board grabble.Board
	test_utils.LoadJSONFixture(t, "testdata/board.fixture", &board)

	var expectedTransposedBoard grabble.Board

	transposedBoard := board.TransposeBoard()
	test_utils.GetGoldenFileJSON(t, transposedBoard, &expectedTransposedBoard, "transposed_board", true)

	assert.Equal(t, expectedTransposedBoard, transposedBoard)

	board[5][7].Letter = 'c'
	assert.Equal(t, board[5][7], transposedBoard[7][5])

}
