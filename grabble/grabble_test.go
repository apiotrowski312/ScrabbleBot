package grabble_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/apiotrowski312/goldtest"
	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func Test_PlaceWord(t *testing.T) {
	type round struct {
		word       string
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
					"WORDS",
					[2]int{7, 7},
					true,
					false,
				},
			},
			"fresh_game.fixture",
		},
		{
			"One round, with blank",
			[]round{
				{
					"WORdS",
					[2]int{7, 7},
					true,
					false,
				},
			},
			"fresh_game.fixture",
		},
		{
			"Two rounds with missing word in dictionary",
			[]round{
				{
					"WORD",
					[2]int{7, 7},
					true,
					false,
				},
				{
					"TESTS",
					[2]int{3, 11},
					false,
					true,
				},
			},
			"fresh_game.fixture",
		},
		{
			"Two rounds",
			[]round{
				{
					"WORD",
					[2]int{7, 7},
					true,
					false,
				},
				{
					"WEST",
					[2]int{7, 7},
					false,
					false,
				},
			},
			"fresh_game.fixture",
		},
		{
			"Second round on error",
			[]round{
				{
					"WORDS",
					[2]int{7, 7},
					true,
					false,
				},
				{
					"WORDSX",
					[2]int{7, 7},
					true,
					true,
				},
			},
			"fresh_game.fixture",
		},
		{
			"No letters to place",
			[]round{
				{
					"",
					[2]int{7, 7},
					true,
					true,
				},
			},
			"fresh_game.fixture",
		},
		{
			"Check if winner is correct",
			[]round{
				{
					"SOS",
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
			test_utils.LoadJSONFixture(t, "../fixtures/"+c.fixture, &game)

			for i, r := range c.rounds {
				err := game.PlaceWord(r.word, r.startPos, r.horizontal)
				assert.Equal(t, r.err, err != nil, fmt.Sprintf("Round: %v\nword: %v\nError: %v", i, r.word, err))
			}
			bytes, _ := json.MarshalIndent(game, "", "\t")
			expectedBytes, err := goldtest.GetGoldenFile(bytes, "testdata/"+c.name)
			if err != nil {
				t.Fatal(err)
			}
			json.Unmarshal(expectedBytes, &expectedGame)

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
			test_utils.LoadJSONFixture(t, "../fixtures/fresh_game.fixture", &game)

			game.PassTurn()
			assert.Equal(t, c.round, game.Stats.CurrentRound)
		})
	}
}
