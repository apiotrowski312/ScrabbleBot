package grabble_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

// TODO: Test vertical stuff maybe
func Test_PickBectWord(t *testing.T) {
	type expectedWord struct {
		word       string
		points     int
		horizontal bool
		cords      [2]int
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
			expectedWord{points: 13, cords: [2]int{7, 3}, word: "WORDS", horizontal: true},
			"testdata/game.fixture",
			"",
		},
		{
			"Get best word - 2",
			expectedWord{points: 17, cords: [2]int{7, 2}, word: "SHROWD", horizontal: true},
			"testdata/game.fixture",
			"../exampleData/collins_official_scrabble_2019.txt",
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

			word := game.PickBestWord()
			assert.Equal(t, c.expectedWord.cords, word.Cords)
			assert.Equal(t, c.expectedWord.horizontal, word.Horizontal)
			assert.Equal(t, c.expectedWord.points, word.Points)
			assert.Equal(t, c.expectedWord.word, word.Word)
		})
	}
}
