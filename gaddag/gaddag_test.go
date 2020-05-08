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
		nameCase string
		hook     rune
		letters  []rune
		words    []string
	}

	t.Run("Small dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/tiny_english.txt")

		cases := []testCase{
			{"Simple test example", 'w', []rune("ord"), []string{"w.ord"}},
			{"Multiple matches", 'w', []rune("ordsk"), []string{"w.ord", "w.ords", "w.ork"}},
			{"Multiple matches from inside", 'r', []rune("orwdsk"), []string{"row.d", "row.ds", "row.k"}},
			{"Single letters", 'b', []rune("ooks"), []string{"b.ook", "b.ooks"}},
			{"O inside hook", 'o', []rune("boooks"), []string{"ob.ok", "oob.k", "oob.ks", "ob.oks"}},
		}

		for _, c := range cases {
			t.Run(c.nameCase, func(t *testing.T) {
				words := gaddagRoot.FindAllWords(c.hook, c.letters)
				assert.ElementsMatch(t, c.words, words)
			})
		}
	})

	t.Run("Full dictionary", func(t *testing.T) {
		gaddagRoot, _ := gaddag.CreateGraph("../exampleData/collins_official_scrabble_2019.txt")

		cases := []testCase{
			{"Simple test example", 'w', []rune("ord"), []string{"w.ord", "wo.", "w.o", "wor.", "word.", "wod."}},
			{"Long word match", 'z', []rune("incographer"), []string{"zihc.", "zinagro.er", "zorra.", "z.aire", "zipac.", "z.inger", "z.e", "z.echin", "zihp.", "z.oecia", "zirohpa.er", "zirg.e", "zir.", "zir.a", "z.ingaro", "zc.ar", "z.o", "zoc.ie", "zaec.ing", "za.onic", "zoc.ing", "zac.", "zag.er", "zarc.e", "zinoga.e", "za.on", "z.incograph", "zoc.", "zan.e", "zag.ier", "zarg.e", "zar.oring", "zahg.i", "z.inc", "znap.er", "zorc.e", "zer.", "z.one", "z.oea", "zah.er", "z.ep", "zirp.er", "zihp.og", "zoc.e", "zan.i", "za.o", "zirg.", "zoc.en", "z.ingare", "z.ine", "z.ero", "zineg.ah", "zorc.ier", "z.ag", "z.arnich", "zah.e", "z.arnec", "zi.ar", "zac.h", "zeipa.on", "z.inco", "z.ip", "z.oner", "z.ein", "zingoc.er", "z.ing", "z.oic", "z.anier", "zinagro.e", "zip.e", "zoc.ier", "zarg.er", "zar.ing", "z.ircon", "z.en", "zinga.e", "zah.ier", "za.ine", "z.ori", "z.ho", "zerp.", "znig.o", "zarc.ier", "zar.or", "zar.er", "zah.ing", "za.oic", "z.ari", "zingoc.e", "zirohpa.e", "zarg.ier", "z.ona", "z.oa", "z.ap", "zipe.oa", "zop.", "zag.on", "zarc.ing", "za.ione", "z.incographer", "z.onae", "zih.en", "zar.e", "zaep.ing", "zeip.o", "zorc.er", "z.in", "z.a", "z.ig", "zirga.e", "zehc.", "zinopac.e", "zirp.e", "zipe.oan", "zan.ir", "zag.e", "z.igan", "zihr.ocarp", "znor.", "znor.er", "z.eroing", "z.ea"}},
		}

		for _, c := range cases {
			t.Run(c.nameCase, func(t *testing.T) {
				words := gaddagRoot.FindAllWords(c.hook, c.letters)
				assert.ElementsMatch(t, c.words, words)
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
	}

	cases := []testCase{
		{"12 letters", 'z', []rune("incographer")},
		{"5 letters", 'w', []rune("odrs")},
		{"15 letters", 'o', []rune("icardehartetis")},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gaddagRoot.FindAllWords(tc.hook, tc.letters)

			}

		})
	}
}
