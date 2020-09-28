package grabble_test

import (
	"fmt"
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
		name          string
		expectedWords []expectedWord
		fixture       string
		dict          string
	}
	test := []testCase{
		{
			"Get best word",
			[]expectedWord{
				{points: 26, word: "WORDS", ratio: 20.76338028169014},
				{points: 26, word: "WORDS", ratio: 20.76338028169014},
				{points: 20, word: "WORDS", ratio: 15.971830985915492},
				{points: 20, word: "WORDS", ratio: 15.971830985915492},
				{points: 24, word: "WOrDS", ratio: 15.27887323943662},
				{points: 24, word: "WOrDS", ratio: 15.27887323943662},
			},
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

			words := game.PickBestWord(len(c.expectedWords))

			for i, word := range words {
				errMessage := fmt.Sprintf("Wrong word is %v, index %v", word.Word, i)
				assert.Equal(t, c.expectedWords[i].word, word.Word, errMessage)
				assert.Equal(t, c.expectedWords[i].points, word.Points, errMessage)
				assert.Equal(t, c.expectedWords[i].ratio, word.Ratio, errMessage)
			}
		})
	}
}
