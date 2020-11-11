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
			{"1 blank", 0, []rune{'W', rune(0), rune(0), rune(0)}, []rune("OR_")},
			{"4 blanks", 0, []rune{'W', rune(0), rune(0), rune(0), rune(0)}, []rune("____")},
			{"Blank as a hook", 0, []rune{'w', rune(0), rune(0), rune(0), rune(0)}, []rune("ORD")},
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
			{"Debug wrong matches for MAGOT, blank letters", 3, []rune{rune(0), rune(0), rune(0), 'm', 'a', 'g', 'o', 't', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("TEICIEI")},
			{"Look for words with blank", 5, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), 'W'}, []rune("ORD_")},
			{"7 letters with existing board and blank", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'B', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WS_ARED")},
			{"7 letters with existing board and 3 blanks", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'B', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WS_A__D")},
			{"7 letters with existing board and 3 blanks. Blank as hook", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'b', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WS_A__D")},
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

func Test_GetNextPermutation(t *testing.T) {
	type testCase struct {
		name                string
		permutation         []rune
		expectedPermutation []rune
		err                 bool
	}
	cases := []testCase{
		{"One letter permutation", []rune("a"), []rune("b"), false},
		{"One letter permutation", []rune("z"), []rune{}, true},
		{"Three letter permutation", []rune("aaa"), []rune("baa"), false},
		{"Three letter permutation 1", []rune("zaa"), []rune("bba"), false},
		{"Three letter permutation 2", []rune("zza"), []rune("bbb"), false},
		{"Three letter permutation 3", []rune("cza"), []rune("dza"), false},
		{"Three letter permutation 4", []rune("azz"), []rune("bzz"), false},
		{"Three letter permutation with error", []rune("zzz"), []rune{}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			permutation, err := gaddag.GetNextPermutation(c.permutation)
			assert.Equal(t, c.expectedPermutation, permutation)
			assert.Equal(t, c.err, err != nil)
		})
	}
}

func Benchmark_CreateGraph(b *testing.B) {
	type testCase struct {
		name string
		dict string
	}

	cases := []testCase{
		{"tiny", "../fixtures/tiny_english.txt"},
		{"100 words dict", "../fixtures/100_english.txt"},
		{"1k words dict", "../fixtures/1k_english.txt"},
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
		{"5 letters + 2 blanks", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AILNR__")},
		{"6 letters + 2 blanks", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AEILNR__")},
		{"8 letters", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'Z', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("AEILNRST")},
		{"7 letters with existring board", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSARED")},
		{"7 letters with existring board. 1 blanks", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSARE_")},
		{"7 letters with existring board. 2 blanks", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSAR__")},
		{"7 letters with existring board. 3 blanks", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSSA___")},
		{"7 letters with existring board. 4 blanks", 8, []rune{'D', 'O', 'W', 'N', rune(0), rune(0), rune(0), rune(0), 'N', rune(0), rune(0), rune(0), 'E', 'S', rune(0)}, []rune("WSS____")},
		{"15 letters. 1 blanks", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'O', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ICARDEHARTETI_")},
		{"15 letters. 1 blanks. Hook max to the right", 15, []rune{rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), 'O'}, []rune("ICARDEHARTETI_")},
		{"15 letters. 1 blanks. Hook max to the left", 0, []rune{'A', rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0), rune(0)}, []rune("ICARDEHARTETI_")},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(c.hookIndex, c.row, c.letters)

			}

		})
	}
}

func Test_SaveToFile(t *testing.T) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/tiny_english.txt")
	gaddagRoot.SaveToFile("../fixtures/test_dict.gaddag")
}

func Test_LoadFromFile(t *testing.T) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/tiny_english.txt")
	gaddagRoot.SaveToFile("../fixtures/test_dict.gaddag")
	gdRoot := gaddag.LoadFromFile("../fixtures/test_dict.gaddag")

	assert.Equal(t, gaddagRoot, &gdRoot)
}

func Benchmark_SaveToFile(b *testing.B) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")

	b.Run(b.Name(), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gaddagRoot.SaveToFile("../fixtures/test_dict.gaddag")
		}
	})
}

func Benchmark_LoadFromFile(b *testing.B) {
	gaddagRoot, _ := gaddag.CreateGraph("../fixtures/collins_official_scrabble_2019.txt")
	gaddagRoot.SaveToFile("../fixtures/test_dict.gaddag")

	b.Run(b.Name(), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			gaddag.LoadFromFile("../fixtures/test_dict.gaddag")
		}
	})
}
