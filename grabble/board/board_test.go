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
				{'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0)},
				{'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0)},
				{'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0)}, // 3
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)},
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)},
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, // 6
				{'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0)},
				{'W', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 's', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'W'},
				{'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0)}, // 9
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)},
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)},
				{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, // 12
				{'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0)},
				{'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0), 'l', 'L', rune(0)},
				{'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0), 'W', 'w', 'W', rune(0), rune(0)}, // 15
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
	board[5][7].Letter = 'C'
	assert.Equal(t, board[5][7], transposedBoard[7][5])

	transposedBoard[5][7].Letter = 'A'
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
			[]string{"WORDS"},
			[][2]int{
				{0, 0},
			},
			true,
		},
		{
			"Place many word",
			[]string{
				"WORDS", "WORDS", "WORDS", "WORDS", "WORDS", "WORDS",
			},
			[][2]int{
				{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
			},
			true,
		},
		{
			"Simple vertically",
			[]string{
				"WORDS",
			},
			[][2]int{
				{0, 0},
			},
			false,
		},
		{
			"Words vertically",
			[]string{
				"WORDS",
				"WORDS",
				"WORDS",
			},
			[][2]int{
				{0, 0}, {0, 1}, {2, 2},
			},
			false,
		},
		{
			"Put blank",
			[]string{
				"WOrdS",
			},
			[][2]int{
				{0, 0},
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
				board.PlaceWord(word, c.startPos[i], "none", 0, c.horizontal)
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
			"WORDS",
			[2]int{7, 7},
			true,
			true,
			[]rune{'W', 'O', 'R', 'D', 'S'},
			"empty_board.fixture",
		},
		{
			"Wrong place - no starting point",
			"WORDS",
			[2]int{0, 0},
			true,
			false,
			[]rune{},
			"empty_board.fixture",
		},
		{
			"Proper vertical with hook",
			"WORDS",
			[2]int{6, 8},
			false,
			true,
			[]rune{'W', 'R', 'D', 'S'},
			"board_with_starting_word.fixture",
		},
		{
			"hook on left",
			"WORDS",
			[2]int{6, 12},
			false,
			true,
			[]rune{'W', 'O', 'R', 'D', 'S'},
			"board_with_starting_word.fixture",
		},
		{
			"No hook",
			"TESTUJ",
			[2]int{6, 13},
			false,
			false,
			[]rune{},
			"board_with_starting_word.fixture",
		},
		{
			"No hook and to long word",
			"TESTUJ",
			[2]int{6, 13},
			true,
			false,
			[]rune{},
			"board_with_starting_word.fixture",
		},
		{
			"hook on left, blank used",
			"WOrDS",
			[2]int{6, 12},
			false,
			true,
			[]rune{'W', 'O', 'r', 'D', 'S'},
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
			"WORDS",
			[2]int{7, 7},
			true,
			[]string{"WORDS"},
			[]string{"s\x00\x00\x00\x00"},
			"empty_board.fixture",
		},
		{
			"Proper vertical with hook",
			"WORDS",
			[2]int{6, 8},
			false,
			[]string{"WORDS"},
			[]string{"\x00\x00\x00\x00\x00"},
			"board_with_starting_word.fixture",
		},
		{
			"hook on left",
			"TEST",
			[2]int{6, 12},
			false,
			[]string{"TEST", "WORDSE"},
			[]string{"l\x00l\x00", "\x00\x00\x00\x00\x00\x00"},
			"board_with_starting_word.fixture",
		},
		{
			"Add x to word",
			"WORDSX",
			[2]int{7, 7},
			true,
			[]string{"WORDSX"},
			[]string{"\x00\x00\x00\x00\x00\x00"},
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
