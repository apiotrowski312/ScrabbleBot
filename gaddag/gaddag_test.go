package gaddag_test

import (
	"errors"
	"flag"
	"fmt"
	"sort"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_CreateGraph(t *testing.T) {
	gaddagRoot, err := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	var expected gaddag.Node
	test_utils.GetGoldenFileJSON(t, gaddagRoot, &expected, t.Name(), *update)

	assert.Equal(t, err, nil)
	assert.Equal(t, &expected, gaddagRoot)
}

func Test_IsWordValid(t *testing.T) {
	gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

	type testCase struct {
		word    string
		errWord string
	}
	testWord := []testCase{
		{"WORD", ""},
		{"W.OR", "WOR"},
		{"W.ORD", ""},
		{"W.ORDX", "WORDX"},
		{"W.ORTHLESS", ""},
		{"W.ORTZHLESS", "WORTZHLESS"},
		{"OB.SS", ""},
		{"SSOB.", ""},
		{"BOSS", ""},
	}

	for _, word := range testWord {
		t.Run(word.word, func(t *testing.T) {
			isOk, err := gaddagRoot.IsWordValid(word.word)
			if len(word.errWord) > 0 {
				assert.Equal(t, false, isOk)
				assert.Equal(t, errors.New(fmt.Sprintf("Word %v is not in dictionary", word.errWord)), err)
			} else {
				assert.Equal(t, true, isOk)
				assert.Equal(t, nil, err)
			}
		})
	}
}

func Test_FindAllWords(t *testing.T) {
	type testCase struct {
		name      string
		hookIndex int
		row       []rune
		letters   []rune
	}

	t.Run("Small dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

		cases := []testCase{
			{"word to right", 0, []rune{'W', rune(0), rune(0), rune(0)}, []rune("ORD")},
			{"word to left", 3, []rune{rune(0), rune(0), rune(0), 'D'}, []rune("ORW")},
			{"Multiple matches", 0, []rune{'W', rune(0), rune(0), rune(0), rune(0)}, []rune("ORDSK")},
			{"Multiple matches from inside", 3, []rune{rune(0), rune(0), rune(0), 'R', rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ORWDSK")},
			{"Single letters", 1, []rune{rune(0), 'B', rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("OOKS")},
			{"O inside hook", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'O', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("BOOOKS")},
			{"with exisitng letters on left", 3, []rune{'W', rune(0), rune(0), 'S'}, []rune("OR")},
			{"with exisitng letters on right", 3, []rune{rune(0), rune(0), rune(0), 'D', 'S'}, []rune("WOR")},
			{"with exisitng letters combo", 3, []rune{'W', rune(0), rune(0), 'D', 'S'}, []rune("OR")},
			{"hook is beetwen other letters", 3, []rune{'W', rune(0), 'R', 'D', 'S'}, []rune("OR")},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var expectedWords []string
				words := gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)
				sort.Strings(words)
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "Small_dictionary/"+c.name, *update)

				assert.Equal(t, expectedWords, words)
			})
		}
	})

	t.Run("Full dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")

		cases := []testCase{
			{"word from left", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'W'}, []rune("ORD")},
			{"Long word match", 2, []rune{rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0)}, []rune("INCOGRAPHER")},
			{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AEILNRST")},
			{"Long word with some letters on board", 1, []rune{rune(0), 'Z', 'I', 'N', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("INCOGRAPHER")},
			{"7 letters with existring board", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'B', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSARED")},
			// FIXME: For some reason, R is duplicated in output 0_0
			{"Test with double R", 7, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'R', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("WGRESEA")},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var expectedWords []string
				words := gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)
				sort.Strings(words)
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "Full_dictionary/"+c.name, *update)

				assert.Equal(t, expectedWords, words)
			})
		}
	})

}

func Benchmark_CreateGraph(b *testing.B) {
	type testCase struct {
		name string
		dict string
	}

	cases := []testCase{
		{"2k words dict", "../exampleData/2k_english.txt"},
		{"20k words dict", "../exampleData/20k_english.txt"},
		{"280k words dict", "../exampleData/collins_official_scrabble_2019.txt"},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddag.CreateGraph(tc.dict)
			}

		})
	}
}

func Benchmark_FindAllWords(b *testing.B) {
	gaddagRoot, _ := gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
	type testCase struct {
		name      string
		hookIndex int
		row       []rune
		letters   []rune
	}

	cases := []testCase{
		{"5 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'W', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("odrs")},
		{"12 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("incographer")},
		{"15 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'o', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("icardehartetis")},
		{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("aeilnrst")},
		{"7 letters with existring board", 8, []rune{'d', 'o', 'W', 'n', rune(0), rune(0), rune(0), rune(0), 'n', rune(0), rune(0), rune(0), 'e', 's', rune(0)}, []rune("wssared")},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)

			}

		})
	}
}
