package gaddag_test

import (
	"errors"
	"flag"
	"fmt"
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
		name    string
		hook    rune
		letters []rune
		left    int
		right   int
	}

	t.Run("Small dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

		cases := []testCase{
			{"Simple test example from left", 'w', []rune("ord"), 0, 3},
			{"Simple test example from right", 'd', []rune("orw"), 5, 0},
			{"Multiple matches", 'w', []rune("ordsk"), 0, 4},
			{"Multiple matches from inside", 'r', []rune("orwdsk"), 3, 5},
			{"Single letters", 'b', []rune("ooks"), 1, 5},
			{"O inside hook", 'o', []rune("boooks"), 5, 13},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var expectedWords []string
				words := gaddagRoot.FindAllWords(c.hook, c.letters, c.left, c.right)
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "Small_dictionary/"+c.name, *update)

				assert.ElementsMatch(t, expectedWords, words)
			})
		}
	})

	t.Run("Full dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")

		cases := []testCase{
			{"Simple test example", 'w', []rune("ord"), 5, 0},
			{"Long word match", 'z', []rune("incographer"), 2, 4},
			{"8 letters", 'z', []rune("aeilnrst"), 15, 15},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var expectedWords []string
				words := gaddagRoot.FindAllWords(c.hook, c.letters, c.left, c.right)
				test_utils.GetGoldenFileJSON(t, words, &expectedWords, "Full_dictionary/"+c.name, *update)

				assert.ElementsMatch(t, expectedWords, words)
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
		name    string
		hook    rune
		letters []rune
		left    int
		right   int
	}

	cases := []testCase{
		{"5 letters", 'w', []rune("odrs"), 15, 15},
		{"12 letters", 'z', []rune("incographer"), 15, 15},
		{"15 letters", 'o', []rune("icardehartetis"), 15, 15},
		{"8 letters", 'z', []rune("aeilnrst"), 15, 15},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(tc.hook, tc.letters, tc.left, tc.right)

			}

		})
	}
}
