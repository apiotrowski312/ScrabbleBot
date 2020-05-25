package grabble_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func Test_CreateBag(t *testing.T) {
	type testCase struct {
		name  string
		tiles []rune
	}
	test := []testCase{
		{
			"Bag of tiles",
			[]rune("abcdefghijklmnoprstuw"),
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedBag grabble.Bag
			bag := grabble.CreateBag(c.tiles)
			test_utils.GetGoldenFileJSON(t, bag, &expectedBag, c.name, *update)

			assert.Equal(t, expectedBag, bag)

		})
	}
}

func Test_CreateLettersPoint(t *testing.T) {
	type testCase struct {
		name string
		lp   map[rune]int
	}
	test := []testCase{
		{
			"Letter values",
			map[rune]int{
				'a': 5,
				'b': 10,
			},
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedLettersPoint grabble.LettersPoint
			lp := grabble.CreateLettersPoint(c.lp)
			test_utils.GetGoldenFileJSON(t, lp, &expectedLettersPoint, c.name, *update)

			assert.Equal(t, expectedLettersPoint, lp)

		})
	}
}

func Test_DrawLetters(t *testing.T) {
	type testCase struct {
		name         string
		tiles        []rune
		drawXLetters int
	}
	test := []testCase{
		{
			"Bag of tiles",
			[]rune("abcdefghijklmnoprstuw"),
			5,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			bag := grabble.CreateBag(c.tiles)
			drawedLetters := bag.DrawLetters(c.drawXLetters)
			assert.Equal(t, c.drawXLetters, len(drawedLetters))
			assert.Equal(t, len(c.tiles)-c.drawXLetters, len(bag))
		})
	}
}
