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

func Test_DrawTiles(t *testing.T) {
	var tb letters.TileBag
	test_utils.LoadJSONFixture(t, "testdata/tilebag.fixture", &tb)

	tiles := tb.DrawTiles(8)
	assert.Equal(t, 8, len(tiles))
	assert.Equal(t, 90, len(tb))

	tiles = tb.DrawTiles(80)
	assert.Equal(t, 80, len(tiles))
	assert.Equal(t, 10, len(tb))

	tiles = tb.DrawTiles(1)
	assert.Equal(t, 1, len(tiles))
	assert.Equal(t, 9, len(tb))

	tiles = tb.DrawTiles(13)
	assert.Equal(t, 9, len(tiles))
	assert.Equal(t, 0, len(tb))

	tiles = tb.DrawTiles(1)
	assert.Equal(t, 0, len(tiles))
	assert.Equal(t, 0, len(tb))
}
