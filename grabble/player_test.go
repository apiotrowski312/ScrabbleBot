package grabble_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

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
			[]rune("abcdefgh"),
			[]rune("abc"),
			[]rune("ijk"),
			false,
		},
		{
			"Test with multiple sameletters",
			[]rune("aabccdde"),
			[]rune("abc"),
			[]rune("iak"),
			false,
		},
		{
			"Test with letter that not exist in rack",
			[]rune("abcdefgh"),
			[]rune("xyz"),
			[]rune("ijk"),
			true,
		},
	}

	for _, c := range test {
		t.Run(c.name, func(t *testing.T) {
			var expectedPlayer grabble.Player
			player := grabble.Player{Rack: c.rack}

			err := player.UpdateRack(c.toRemove, c.toAdd)
			test_utils.GetGoldenFileJSON(t, player, &expectedPlayer, c.name, true)

			assert.Equal(t, expectedPlayer, player)
			assert.Equal(t, c.err, err != nil)
		})
	}
}
