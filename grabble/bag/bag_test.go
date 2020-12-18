package bag_test

import (
	"testing"

	"github.com/apiotrowski312/goldtest"
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
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
			[]rune("ABCDEFGHIJKLMNOPRSTUWXYZ"),
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			bag := bag.CreateBag(c.tiles)
			goldtest.AssertJSON(t, bag, "testdata/"+c.name)
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
				'A': 1,
				'B': 1,
				'C': 1,
				'D': 2,
				'E': 2,
				'F': 2,
				'G': 3,
				'H': 3,
				'I': 3,
				'J': 4,
				'K': 4,
				'L': 4,
				'M': 1,
				'N': 1,
				'O': 1,
				'P': 2,
				'R': 2,
				'S': 2,
				'T': 3,
				'U': 3,
				'W': 3,
				'X': 4,
				'Y': 4,
				'Z': 4,
			},
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			lp := bag.CreateLettersPoint(c.lp)

			goldtest.AssertJSON(t, lp, "testdata/"+c.name)
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
			[]rune("ABCDEFGHIJKLMNOPRSTUW"),
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
			[]rune("ABCDEFGHIJKLMNOPRSTUW"),
			[]rune("QWE"),
		},
		{
			"Empty bag",
			[]rune(""),
			[]rune("QAWSEDRF"),
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

func Test_GetPoints(t *testing.T) {
	type testCase struct {
		name    string
		words   []string
		bonuses []string
		points  int
	}
	test := []testCase{
		{
			"One word",
			[]string{"WORDS"},
			[]string{"0W000"},
			30,
		},
		{
			"Starting Point",
			[]string{"WORDS"},
			[]string{"0s000"},
			20,
		},
		{
			"Blank test, word bonus",
			[]string{"WoRDS"},
			[]string{"0w000"},
			18,
		},
		{
			"Blank test, letter bonus",
			[]string{"WoRDS"},
			[]string{"0L000"},
			9,
		},
		{
			"Multiple word",
			[]string{"WORDS", "TEST", "BILING"},
			[]string{"0W0W0", "00L0", "00W000"},
			149,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var lv bag.LettersPoint
			test_utils.LoadJSONFixture(t, "../../fixtures/letter_values.fixture", &lv)

			points := lv.GetPoints(c.words, c.bonuses)
			assert.Equal(t, c.points, points)
		})
	}
}
