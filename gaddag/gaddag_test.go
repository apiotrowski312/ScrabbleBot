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
		{"word", ""},
		{"w.or", "wor"},
		{"w.ord", ""},
		{"w.ordX", "wordX"},
		{"w.orthless", ""},
		{"w.ortzhless", "wortzhless"},
		{"ob.ss", ""},
		{"ssob.", ""},
		{"boss", ""},
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
			{"word to right", 0, []rune{'w', rune(0), rune(0), rune(0)}, []rune("ord")},
			{"word to left", 3, []rune{rune(0), rune(0), rune(0), 'd'}, []rune("orw")},
			{"Multiple matches", 0, []rune{'w', rune(0), rune(0), rune(0), rune(0)}, []rune("ordsk")},
			{"Multiple matches from inside", 3, []rune{rune(0), rune(0), rune(0), 'r', rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("orwdsk")},
			{"Single letters", 1, []rune{rune(0), 'b', rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ooks")},
			{"O inside hook", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'o', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("boooks")},
			{"with exisitng letters on left", 3, []rune{'w', rune(0), rune(0), 'd'}, []rune("or")},
			{"with exisitng letters on right", 3, []rune{rune(0), rune(0), rune(0), 'd', 's'}, []rune("wor")},
			{"with exisitng letters combo", 3, []rune{'w', rune(0), rune(0), 'd', 's'}, []rune("or")},
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
			{"word from left", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'w'}, []rune("ord")},
			{"Long word match", 2, []rune{rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0)}, []rune("incographer")},
			{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("aeilnrst")},
			{"Long word with some letters on board", 1, []rune{rune(0), 'z', 'i', 'n', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("incographer")},
			{"7 letters with existring board", 8, []rune{'d', 'o', 'w', 'n', rune(0), rune(0), rune(0), rune(0), 'n', rune(0), rune(0), rune(0), 'e', 's', rune(0)}, []rune("wssared")},
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
		{"5 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'w', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("odrs")},
		{"12 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("incographer")},
		{"15 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'o', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("icardehartetis")},
		{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("aeilnrst")},
		{"7 letters with existring board", 8, []rune{'d', 'o', 'w', 'n', rune(0), rune(0), rune(0), rune(0), 'n', rune(0), rune(0), rune(0), 'e', 's', rune(0)}, []rune("wssared")},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)

			}

		})
	}
}
