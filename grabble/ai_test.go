package grabble_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func Test_PickBectWord(t *testing.T) {

	type testCase struct {
		name         string
		fixture      string
		expectedWord string
	}
	test := []testCase{
		{
			"Test pass round",
			"testdata/game.fixture",
			"food",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			test_utils.LoadJSONFixture(t, c.fixture, &game)

			word := game.PickBestWord()
			assert.Equal(t, c.expectedWord, word)
		})
	}
}

func Test_PickBectWord_BigData(t *testing.T) {

	type testCase struct {
		name         string
		fixture      string
		expectedWord string
	}
	test := []testCase{
		{
			"Test pass round",
			"testdata/game.fixture",
			"food",
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var game grabble.Grabble
			test_utils.LoadJSONFixture(t, c.fixture, &game)
			gaddagRoot, _ := gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
			game.Dict = *gaddagRoot

			word := game.PickBestWord()
			assert.Equal(t, c.expectedWord, word)
		})
	}
}
