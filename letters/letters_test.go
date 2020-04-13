package letters_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/letters"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_LoadTiles(t *testing.T) {
	tb, lv, err := letters.LoadTilesFromFile("../exampleData/english_tiles.csv")

	expectedTbBytes := test_utils.GetGoldenFileJSON(t, tb, t.Name()+"_tiles", *update)
	expectedLvBytes := test_utils.GetGoldenFileJSON(t, lv, t.Name()+"_letters", *update)

	var expectedLv letters.LetterValue
	var expectedTb letters.TileBag

	test_utils.BytesToStruct(t, expectedLvBytes, &expectedLv)
	test_utils.BytesToStruct(t, expectedTbBytes, &expectedTb)

	assert.Equal(t, nil, err)
	assert.Equal(t, &expectedTb, tb)
	assert.Equal(t, &expectedLv, lv)

}

func Test_CountPoints(t *testing.T) {
	var lv letters.LetterValue
	test_utils.LoadJSONFixture(t, "testdata/letters_values.fixture", &lv)

	words := []string{"socks"}
	tiles := []string{"0L0L0"}

	points := lv.CountPoints(words, tiles)
	assert.Equal(t, 23, points)

	words = []string{"socks", "sdog."}
	tiles = []string{"0L0L0", "0"}

	points = lv.CountPoints(words, tiles)
	assert.Equal(t, 29, points)

	words = []string{"socks"}
	tiles = []string{"WLslw"}

	points = lv.CountPoints(words, tiles)
	assert.Equal(t, 108, points)
}
