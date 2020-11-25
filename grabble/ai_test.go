package grabble_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func Test_PickBectWord(t *testing.T) {
	type expectedWord struct {
		word   string
		points int
		ratio  float64
	}
	type testCase struct {
		name         string
		expectedWord expectedWord
		fixture      string
		dict         string
	}
	test := []testCase{
		{
			"Get best word",
			expectedWord{points: 26, word: "WORDS", ratio: 20.76338028169014},
			"../fixtures/fresh_game.fixture",
			"",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			test_utils.LoadJSONFixture(t, c.fixture, &game)
			if c.dict != "" {
				gaddagRoot, _ := gaddag.CreateGraph(c.dict)
				game.Dict = *gaddagRoot
			}

			word, _ := game.PickBestWord()

			assert.Equal(t, c.expectedWord.word, word.Word)
			assert.Equal(t, c.expectedWord.points, word.Points)
			assert.Equal(t, c.expectedWord.ratio, word.Ratio)
		})
	}
}
