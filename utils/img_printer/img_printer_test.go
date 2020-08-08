package img_printer_test

import (
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/utils/img_printer"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func TestNewSceneReturnsANewScene(t *testing.T) {
	var game grabble.Grabble
	test_utils.LoadJSONFixture(t, "../../grabble/testdata/endgame.fixture", &game)

	img_printer.PrintScreenBoard(game, "test.png")
}
