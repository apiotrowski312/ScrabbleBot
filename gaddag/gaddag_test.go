package gaddag_test

import (
	"errors"
	"flag"
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
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

// func Test_IsWordValid(t *testing.T) {
// 	gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

// }

func Benchmark_CreateGraph_2kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/2k_english.txt")
	}
}

func Benchmark_CreateGraph_20kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/20k_english.txt")
	}
}

func Benchmark_CreateGraph_280kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
	}
}
