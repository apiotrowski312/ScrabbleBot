package grabble

import (
	"github.com/apiotrowski312/scrabbleBot/grabble/bag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
	"github.com/apiotrowski312/scrabbleBot/grabble/player"
)

type grabble struct {
	board         board.Board
	players       []player.Player
	bag           bag.Bag
	lettterPoints bag.LettersPoint
}

// TODO: Place word
func (g *grabble) PlaceWord(word string, letters []rune, startPos [2]int, horizontal bool) {

}

// TODO: Exception when all letters were used (plus 50 points)
