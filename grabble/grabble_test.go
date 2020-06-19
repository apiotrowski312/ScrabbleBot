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
		name   string
		rounds []round
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
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			var expectedGame grabble.Grabble
			test_utils.LoadJSONFixture(t, "testdata/game.fixture", &game)

			for i, r := range c.rounds {
				err := game.PlaceWord(r.word, r.letters, r.startPos, r.horizontal)
				assert.Equal(t, r.err, err != nil, fmt.Sprintf("Round: %v, word: %v", i, r.word))
			}

			test_utils.GetGoldenFileJSON(t, game, &expectedGame, c.name, *update)
			assert.Equal(t, expectedGame, game)
		})
	}
}
