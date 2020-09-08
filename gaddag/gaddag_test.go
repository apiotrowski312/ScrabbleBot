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
	gaddagRoot, err := gaddag.CreateGraph("../fixtures/tiny_english.txt")

	var expected gaddag.Node
	test_utils.GetGoldenFileJSON(t, gaddagRoot, &expected, t.Name(), *update)

	assert.Equal(t, err, nil)
	assert.Equal(t, &expected, gaddagRoot)
}

func Test_IsWordValid(t *testing.T) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/tiny_english.txt")

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
		gaddagRoot, _ := gaddag.CreateGraph("../fixtures/tiny_english.txt")

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
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "old/Small_dictionary/"+c.name, *update)

				assert.Equal(t, expectedWords, words)
			})
		}
	})

	t.Run("Full dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")

		cases := []testCase{
			{"word from left", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'W'}, []rune("ORD")},
			{"Long word match", 2, []rune{rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0)}, []rune("INCOGRAPHER")},
			{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AEILNRST")},
			{"Long word with some letters on board", 1, []rune{rune(0), 'Z', 'I', 'N', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("INCOGRAPHER")},
			{"7 letters with existring board", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'B', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSARED")},
			{"Test for a but with double R", 7, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'R', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("WGESEA")},
			{"Check if DO will be find", 7, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'D', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("VOQFEII")},
			{"Check if LONE will be find", 8, []rune{rune(0), rune(0), rune(0), 'R', rune(0), rune(0), 'H', rune(0), 'L', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("IQEIOCN")},
			{"Debug wrong matches for MAGOT", 3, []rune{rune(0), rune(0), rune(0), 'M', 'A', 'G', 'O', 'T', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("TEICIEI")},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var expectedWords []string
				words := gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)
				sort.Strings(words)
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "old/Full_dictionary/"+c.name, *update)

				assert.Equal(t, expectedWords, words)
			})
		}
	})
}

func Test_FindWords(t *testing.T) {
	type testCase struct {
		name    string
		letters []rune
	}
	cases := []testCase{
		{"simple case", []rune("WORD")},
		{"long case", []rune("WSSARED")},
		{"longest case", []rune("WSSAREDDREFCDSAA")},
		{"Intersing case with extremly hard rack", []rune("QZJXKHW")},
	}

	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var expectedWords []string
			words := gaddagRoot.FindWords(c.letters)
			sort.Strings(words)
			test_utils.GetGoldenFileJSON(t, words, &expectedWords, t.Name()+c.name, *update)

			assert.Equal(t, expectedWords, words)
		})
	}
}

func Benchmark_CreateGraph(b *testing.B) {
	type testCase struct {
		name string
		dict string
	}

	cases := []testCase{
		{"2k words dict", "../fixtures/2k_english.txt"},
		{"20k words dict", "../fixtures/20k_english.txt"},
		{"280k words dict", "../fixtures/collins_official_scrabble_2019.txt"},
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
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")
	type testCase struct {
		name      string
		hookIndex int
		row       []rune
		letters   []rune
	}

	cases := []testCase{
		{"5 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'W', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ODRS")},
		{"12 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("INCOGRAPHER")},
		{"15 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'O', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ICARDEHARTETIS")},
		{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AEILNRST")},
		{"7 letters with existring board", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSARED")},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)

			}

		})
	}
}

func Benchmark_FindWords(b *testing.B) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")
	type testCase struct {
		name    string
		letters []rune
	}

	cases := []testCase{
		{"5 letters", []rune("WODRS")},
		{"7 letters", []rune("WSSARED")},
		{"8 letters", []rune("AEILNRST")},
		{"12 letters", []rune("INCOGRAPHERS")},
		{"15 letters", []rune("WICARDEHARTETIS")},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindWords(c.letters)
			}
		})
	}
}
