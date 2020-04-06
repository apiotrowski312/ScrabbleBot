package game

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

func Test_loadTiles(t *testing.T) {
	tb, lv, err := loadTilesFromFile("../exampleData/english_tiles.csv")

	expectedTbBytes := test_utils.GetGoldenFileJSON(t, tb, t.Name()+"_tiles", *update)
	expectedLvBytes := test_utils.GetGoldenFileJSON(t, lv, t.Name()+"_letters", *update)

	var expectedLv letterValue
	var expectedTb tileBag

	test_utils.BytesToStruct(t, expectedLvBytes, &expectedLv)
	test_utils.BytesToStruct(t, expectedTbBytes, &expectedTb)

	assert.Equal(t, nil, err)
	assert.Equal(t, &expectedTb, tb)
	assert.Equal(t, &expectedLv, lv)

}

func Test_countPoints(t *testing.T) {
	var lv letterValue
	test_utils.LoadJSONFixture(t, "testdata/letters_values.fixture", &lv)

	words := []string{"socks"}
	tiles := []string{"0L0L0"}

	points := lv.countPoints(words, tiles)
	assert.Equal(t, 23, points, "Not working")

	words = []string{"socks", "sdog."}
	tiles = []string{"0L0L0", "0"}

	points = lv.countPoints(words, tiles)
	assert.Equal(t, 29, points, "Not working")

	words = []string{"socks"}
	tiles = []string{"WLslw"}

	points = lv.countPoints(words, tiles)
	assert.Equal(t, 108, points, "Not working")
}
