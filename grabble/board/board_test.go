package board_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble/board"
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
				{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'}, // 3
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}, // 6
				{'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0'},
				{'W', '0', '0', '0', '0', '0', '0', 's', '0', '0', '0', '0', '0', '0', 'W'},
				{'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0'}, // 9
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}, // 12
				{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'},
				{'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0', 'l', 'L', '0'},
				{'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0', 'W', 'w', 'W', '0', '0'}, // 15
			},
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedBoard board.Board
			board := board.CreateBoard(c.template)

			test_utils.GetGoldenFileJSON(t, board, &expectedBoard, c.name, *update)

			assert.Equal(t, &expectedBoard, board)
		})
	}
}

func Test_TransposeBoard(t *testing.T) {
	var expectedTransposedBoard board.Board
	var board board.Board
	test_utils.LoadJSONFixture(t, "../../fixtures/empty_board.fixture", &board)

	transposedBoard := board.TransposeBoard()
	board[5][7].Letter = 'c'
	assert.Equal(t, board[5][7], transposedBoard[7][5])

	transposedBoard[5][7].Letter = 'a'
	assert.Equal(t, transposedBoard[5][7], board[7][5])

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
			var expectedBoard board.Board
			var board board.Board

			test_utils.LoadJSONFixture(t, "../../fixtures/empty_board.fixture", &board)

			for i, word := range c.words {
				board.PlaceWord(word, c.startPos[i], c.horizontal)
			}

			test_utils.GetGoldenFileJSON(t, board, &expectedBoard, c.name, *update)
			assert.Equal(t, expectedBoard, board)
		})
	}
}

func Test_DoesHookExist(t *testing.T) {
	type testCase struct {
		name       string
		word       string
		startPos   [2]int
		horizontal bool
		isOk       bool
		letters    []rune
		fixture    string
	}
	test := []testCase{
		{
			"Word in the middle",
			"words",
			[2]int{7, 7},
			true,
			true,
			[]rune{'w', 'o', 'r', 'd', 's'},
			"empty_board.fixture",
		},
		{
			"Wrong place - no starting point",
			"words",
			[2]int{0, 0},
			true,
			false,
			[]rune{},
			"empty_board.fixture",
		},
		{
			"Proper vertical with hook",
			"words",
			[2]int{6, 8},
			false,
			true,
			[]rune{'w', 'r', 'd', 's'},
			"board_with_starting_word.fixture",
		},
		{
			"hook on left",
			"words",
			[2]int{6, 12},
			false,
			true,
			[]rune{'w', 'o', 'r', 'd', 's'},
			"board_with_starting_word.fixture",
		},
		{
			"No hook",
			"testuj",
			[2]int{6, 13},
			false,
			false,
			[]rune{},
			"board_with_starting_word.fixture",
		},
		{
			"No hook and to long word",
			"testuj",
			[2]int{6, 13},
			true,
			false,
			[]rune{},
			"board_with_starting_word.fixture",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var board board.Board
			test_utils.LoadJSONFixture(t, "../../fixtures/"+c.fixture, &board)
			letters, isOk := board.DoesHookExist(c.word, c.startPos, c.horizontal)

			assert.Equal(t, c.isOk, isOk)
			assert.Equal(t, c.letters, letters)
		})
	}
}

func Test_GetAllWordsAndBonuses(t *testing.T) {
	type testCase struct {
		name       string
		word       string
		startPos   [2]int
		horizontal bool
		words      []string
		bonuses    []string
		fixture    string
	}
	test := []testCase{
		{
			"Word in the middle",
			"words",
			[2]int{7, 7},
			true,
			[]string{"words"},
			[]string{"s0000"},
			"empty_board.fixture",
		},
		{
			"Proper vertical with hook",
			"words",
			[2]int{6, 8},
			false,
			[]string{"words"},
			[]string{"00000"},
			"board_with_starting_word.fixture",
		},
		{
			"hook on left",
			"test",
			[2]int{6, 12},
			false,
			[]string{"test", "wordse"},
			[]string{"l0l0", "000000"},
			"board_with_starting_word.fixture",
		},
		{
			"Add x to word",
			"wordsx",
			[2]int{7, 7},
			true,
			[]string{"wordsx"},
			[]string{"000000"},
			"board_with_starting_word.fixture",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var board board.Board
			test_utils.LoadJSONFixture(t, "../../fixtures/"+c.fixture, &board)
			words, bonuses := board.GetAllWordsAndBonuses(c.word, c.startPos, c.horizontal)

			assert.Equal(t, c.words, words)
			assert.Equal(t, c.bonuses, bonuses)
		})
	}
}
