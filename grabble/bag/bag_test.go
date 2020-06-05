package bag_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

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
			var expectedBag bag.Bag
			bag := bag.CreateBag(c.tiles)
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
			var expectedLettersPoint bag.LettersPoint
			lp := bag.CreateLettersPoint(c.lp)
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
			bag := bag.CreateBag(c.tiles)
			drawedLetters := bag.DrawLetters(c.drawXLetters)
			assert.Equal(t, c.drawXLetters, len(drawedLetters))
			assert.Equal(t, len(c.tiles)-c.drawXLetters, len(bag))
		})
	}
}

func Test_ChangeLetters(t *testing.T) {
	type testCase struct {
		name          string
		tiles         []rune
		changeLetters []rune
	}
	test := []testCase{
		{
			"Change 3 letters",
			[]rune("abcdefghijklmnoprstuw"),
			[]rune("qwe"),
		},
		{
			"Empty bag",
			[]rune(""),
			[]rune("qawsedrf"),
		},
		{
			"Empty bag, changing no letters",
			[]rune(""),
			[]rune(""),
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			bag := bag.CreateBag(c.tiles)
			drawedLetters := bag.ChangeLetters(c.changeLetters)
			assert.Equal(t, len(c.changeLetters), len(drawedLetters))
			assert.Equal(t, len(c.tiles), len(bag))
		})
	}
}
