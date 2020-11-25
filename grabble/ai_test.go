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

func Benchmark_PickBestWord(b *testing.B) {
	type testCase struct {
		name    string
		fixture string
		dict    string
	}
	cases := []testCase{
		{"default strategy - small dict", "../fixtures/fresh_game.fixture", ""},
		{"middle ratio strategy - small dict", "../fixtures/fresh_game_middle_ratio.fixture", ""},
		{"middle ratio strategy", "../fixtures/fresh_game_middle_ratio.fixture", "../fixtures/english_dict.txt"},
		{"default strategy - 0 blank", "../fixtures/fresh_game.fixture", "../fixtures/english_dict.txt"},
		{"default strategy - 2 blank", "../fixtures/fresh_game_2_blank.fixture", "../fixtures/english_dict.txt"},
		{"default strategy - 3 blank", "../fixtures/fresh_game_3_blank.fixture", "../fixtures/english_dict.txt"},
		{"round 20 - default strategy - 0 blank", "../fixtures/20_round_game.fixture", "../fixtures/english_dict.txt"},
		{"round 20 - default strategy - 1 blank", "../fixtures/20_round_game_1_blank.fixture", "../fixtures/english_dict.txt"},
		{"round 20 - default strategy - 2 blank", "../fixtures/20_round_game_2_blank.fixture", "../fixtures/english_dict.txt"},
		{"round 20 - default strategy - 3 blank", "../fixtures/20_round_game_3_blank.fixture", "../fixtures/english_dict.txt"},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				var game grabble.Grabble
				// Fix this nil
				test_utils.LoadJSONFixture(nil, c.fixture, &game)
				if c.dict != "" {
					gaddagRoot, _ := gaddag.CreateGraph(c.dict)
					game.Dict = *gaddagRoot
				}
				b.StartTimer()
				game.PickBestWord()
			}

		})
	}

}
