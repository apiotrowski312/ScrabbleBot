package grabble_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_PlaceWord(t *testing.T) {
	type round struct {
		word       string
		letters    []rune
		startPos   [2]int
		horizontal bool
		err        bool
	}
	type testCase struct {
		name    string
		rounds  []round
		fixture string
	}
	test := []testCase{
		{
			"One round",
			[]round{
				{
					"words",
					[]rune{'w', 'o', 'r', 'd', 's'},
					[2]int{7, 7},
					true,
					false,
				},
			},
			"game.fixture",
		},
		{
			"Three rounds",
			[]round{
				{
					"word",
					[]rune{'w', 'o', 'r', 'd'},
					[2]int{7, 7},
					true,
					false,
				},
				{
					"worthful",
					[]rune{'w', 'r', 't', 'h', 'f', 'u', 'l'},
					[2]int{6, 8},
					false,
					false,
				},
				{
					"sos",
					[]rune{'s', 'o', 's'},
					[2]int{7, 11},
					false,
					false,
				},
			},
			"game.fixture",
		},
		{
			"Second round on error",
			[]round{
				{
					"words",
					[]rune{'w', 'o', 'r', 'd', 's'},
					[2]int{7, 7},
					true,
					false,
				},
				{
					"wordsX",
					[]rune{'X'},
					[2]int{7, 7},
					true,
					true,
				},
			},
			"game.fixture",
		},
		{
			"No letters to place",
			[]round{
				{
					"",
					[]rune{},
					[2]int{7, 7},
					true,
					true,
				},
			},
			"game.fixture",
		},
		{
			"Check if winner is correct",
			[]round{
				{
					"sos",
					[]rune{'s', 'o', 's'},
					[2]int{7, 7},
					true,
					false,
				},
			},
			"endgame.fixture",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			var expectedGame grabble.Grabble
			test_utils.LoadJSONFixture(t, "testdata/"+c.fixture, &game)

			for i, r := range c.rounds {
				err := game.PlaceWord(r.word, r.letters, r.startPos, r.horizontal)
				assert.Equal(t, r.err, err != nil, fmt.Sprintf("Round: %v, word: %v", i, r.word))
			}

			test_utils.GetGoldenFileJSON(t, game, &expectedGame, c.name, *update)
			for i, p := range expectedGame.Players {
				assert.Equal(t, p.Name, game.Players[i].Name)
				assert.Equal(t, p.Points, game.Players[i].Points)
			}
			assert.Equal(t, expectedGame.Board, game.Board)
			assert.Equal(t, expectedGame.Dict, game.Dict)
			assert.Equal(t, expectedGame.LettterPoints, game.LettterPoints)
			assert.Equal(t, expectedGame.Stats, game.Stats)
			assert.Equal(t, expectedGame.RackSize, game.RackSize)
		})
	}
}

func Test_PassTurn(t *testing.T) {

	type testCase struct {
		name  string
		round int
	}
	test := []testCase{
		{
			"Test pass round",
			1,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			test_utils.LoadJSONFixture(t, "testdata/game.fixture", &game)

			game.PassTurn()
			assert.Equal(t, c.round, game.Stats.CurrentRound)
		})
	}
}
