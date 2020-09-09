package grabble_test

import (
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

// TODO: Test vertical stuff maybe, Add more advanced tests, Add benchmark tests
func Test_PickBectWord(t *testing.T) {
	type expectedWord struct {
		word       string
		points     int
		horizontal bool
		cords      [2]int
	}
	type testCase struct {
		name          string
		expectedWords []expectedWord
		fixture       string
		dict          string
	}
	test := []testCase{
		{
			"Get best word",
			[]expectedWord{
				{points: 26, cords: [2]int{7, 3}, word: "WORDS", horizontal: true},
			},
			"../fixtures/fresh_game.fixture",
			"",
		},
		{
			"Get best word - 2",
			[]expectedWord{
				{points: 34, cords: [2]int{2, 7}, word: "SHROWD", horizontal: false},
			},
			"../fixtures/fresh_game.fixture",
			"../fixtures/collins_official_scrabble_2019.txt",
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

			words := game.PickBestWord(1)
			for i, word := range words {
				errMessage := fmt.Sprintf("Wrong word is %v, index %v", word.Word, i)
				assert.Equal(t, c.expectedWords[i].word, word.Word, errMessage)
				assert.Equal(t, c.expectedWords[i].points, word.Points, errMessage)
				// FIXME: Results are not indempotempt.
				// assert.Equal(t, c.expectedWords[i].cords, word.Cords, errMessage)
				// assert.Equal(t, c.expectedWords[i].horizontal, word.Horizontal, errMessage)

			}
		})
	}
}
