package gaddag

import (
	"errors"
	"flag"
	"fmt"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_GaddagAddWord_One(t *testing.T) {
	gd := Node{
		children: map[rune]Node{},
	}
	gd.addWord("word")

	expect := test_utils.GetGoldenFileString(t, fmt.Sprint(gd), t.Name(), *update)
	assert.Equal(t, expect, fmt.Sprint(gd))
}

func Test_GaddagAddWords_Five(t *testing.T) {
	gd := Node{
		children: map[rune]Node{},
	}

	gd.addWord("word")
	gd.addWord("words")
	gd.addWord("work")
	gd.addWord("worthless")
	gd.addWord("worthful")

	expect := test_utils.GetGoldenFileString(t, fmt.Sprint(gd), t.Name(), *update)
	assert.Equal(t, expect, fmt.Sprint(gd))
}

func Test_CreateGraph(t *testing.T) {
	gaddagRoot, err := CreateGraph("../exampleData/tiny_english.txt")
	expect := test_utils.GetGoldenFileString(t, fmt.Sprint(gaddagRoot), t.Name(), *update)

	assert.Equal(t, err, nil)
	assert.Equal(t, expect, fmt.Sprint(gaddagRoot))
}

func Test_IsWordValid(t *testing.T) {
	gaddagRoot, _ := CreateGraph("../exampleData/tiny_english.txt")

	isOk, err := gaddagRoot.IsWordValid("w.or")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word wor is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("w.ord")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.ordX")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word wordX is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("w.orthless")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)
	isOk, err = gaddagRoot.IsWordValid("w.orthleks")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word worthleks is not in dictionary"), err)
	isOk, err = gaddagRoot.IsWordValid("ob.ss")
	assert.Equal(t, true, isOk)
	assert.Equal(t, nil, err)

	isOk, err = gaddagRoot.IsWordValid("notexistingword")
	assert.Equal(t, false, isOk)
	assert.Equal(t, errors.New("Word notexistingword was passed in wrong format (no necessarry dot)"), err)
}

func Benchmark_CreateGraph_5Words(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/tiny_english.txt")
	}
}

func Benchmark_CreateGraph_2kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/2k_english.txt")
	}
}

func Benchmark_CreateGraph_20kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/20k_english.txt")
	}
}

func Benchmark_CreateGraph_280kWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateGraph("../exampleData/collins_official_scrabble_2019.txt")
	}
}
