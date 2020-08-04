package imgPrinter_test

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/apiotrowski312/scrabbleBot/grabble/imgPrinter"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
)

func TestNewSceneReturnsANewScene(t *testing.T) {
	var game grabble.Grabble
	test_utils.LoadJSONFixture(t, "../testdata/game.fixture", &game)

	imgPrinter.PrintScreenBoard(game, "test.png")
}
