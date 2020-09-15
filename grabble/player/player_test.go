package player_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble/player"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_UpdateRack(t *testing.T) {
	type testCase struct {
		name     string
		rack     []rune
		toRemove []rune
		toAdd    []rune
		err      bool
	}
	test := []testCase{
		{
			"Simple test",
			[]rune("ABCDEFGH"),
			[]rune("ABC"),
			[]rune("IJK"),
			false,
		},
		{
			"Test with multiple sameletters",
			[]rune("AABCCDDE"),
			[]rune("ABC"),
			[]rune("IAK"),
			false,
		},
		{
			"Test with letter that not exist in rack",
			[]rune("ABCDEFGH"),
			[]rune("XYZ"),
			[]rune("IJK"),
			true,
		},
		{
			"Test removing blanks",
			[]rune("ABCDE___"),
			[]rune("abc"),
			[]rune("IJK"),
			false,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedPlayer player.Player
			player := player.Player{Rack: c.rack}

			err := player.UpdateRack(c.toRemove, c.toAdd)
			test_utils.GetGoldenFileJSON(t, player, &expectedPlayer, c.name, *update)

			assert.Equal(t, expectedPlayer, player)
			assert.Equal(t, c.err, err != nil)
		})
	}
}

func Test_AreLettersInRack(t *testing.T) {
	type testCase struct {
		name    string
		rack    []rune
		letters []rune
		err     bool
	}
	test := []testCase{
		{"Easy rack", []rune("ABCDEF"), []rune("ABC"), false},
		{"To much A letter", []rune("ABCDEF"), []rune("ABCA"), true},
		{"Check for double letters", []rune("ABBA"), []rune("BABA"), false},
		{"Missing letters in rack", []rune("BA"), []rune("ABBA"), true},
		{"Blanks in rack", []rune("TE__"), []rune("TEst"), false},
		{"One blank only in rack", []rune("TE_"), []rune("TEst"), true},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			player := player.Player{Rack: c.rack}
			err := player.AreLettersInRack(c.letters)

			assert.Equal(t, c.err, err != nil, err)
		})
	}
}
