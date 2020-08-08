package img_printer_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
)

func TestNewSceneReturnsANewScene(t *testing.T) {
	var game grabble.Grabble
	test_utils.LoadJSONFixture(t, "../testdata/endgame.fixture", &game)

	img_printer.PrintScreenBoard(game, "test.png")
}
