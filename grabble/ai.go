package grabble

import (
	"sort"
	"strings"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
)

// TODO: Function with starting word
// FIXME: Better naming for this variable
type gaddagWord struct {
	Points     int
	Cords      [2]int
	Word       string
	Horizontal bool
}

// PickBestWord returns best words for current player and board
func (g *Grabble) PickBestWord(numberOfWords int) []gaddagWord {
	log.Debugf("PickBestWord function called by %s", g.CurrentPlayer().Name)
	log.Debugf("Rack: %s", string(g.CurrentPlayer().Rack))
	rack := g.CurrentPlayer().Rack

	wordsCollection := g.getWordCollection(rack, true)
	wordsCollection = append(wordsCollection, g.getWordCollection(rack, false)...)

	log.Debugf("Found %v words", len(wordsCollection))
	sort.Slice(wordsCollection, func(i, j int) bool {
		return wordsCollection[i].Points > wordsCollection[j].Points
	})

	if len(wordsCollection) < numberOfWords {
		return wordsCollection
	}
	return wordsCollection[:numberOfWords]
}

func (g *Grabble) getWordCollection(rack []rune, horizontal bool) []gaddagWord {
	log.Debugf("getWordCollection called. Horizontal: %v", horizontal)

	board := g.Board
	if !horizontal {
		board = *g.Board.TransposeBoard()
	}

	wordsCollection := []gaddagWord{}

	for x, row := range board {
		rowLetters := board.GetRowOfLetters(x)
		for y := range row {
			words := []string{}
			if board[x][y].Letter != rune(0) && (y == 0 || (y > 0 && board[x][y-1].Letter == rune(0))) {
				log.Debugf("Hook found %v(%v). Horizontal: %v", string(board[x][y].Letter), [2]int{x, y}, horizontal)
				log.Debugf("Row for finding words %v, rack %v, hookIndex %v", rowLetters, rack, y)
				words = g.Dict.FindAllWords(y, rowLetters, rack)
			} else if board[x][y].Bonus == 's' && board[x][y].Letter == rune(0) {
				// HACK: Better and faster option will be create new function in gaddag to look for words without hook
				for i, l := range rack {
					rackForThisItteration := append(append([]rune{}, rack[:i]...), rack[i+1:]...)
					sWords := g.Dict.FindAllWords(y, mockRowForStartingTile(y, l, rowLetters), rackForThisItteration)

					words = append(words, sWords...)
				}
			}
			if len(words) != 0 {
				log.Debugf("There is %v new words before counting points", len(words))
				for _, w := range words {
					normalizedWord, cords := prepareWordAndFixCords(w, [2]int{x, y}, horizontal)
					log.Debugf("Before normalization %v %v", w, [2]int{x, y})
					log.Debugf("After normalization %v %v", normalizedWord, cords)

					letters, letterError := g.validateAndExtractUsedNewLetters(normalizedWord, cords, horizontal)
					if letterError != nil {
						log.Debugf("There was error after validatin word: %v. Error: %v", w, letterError)
						continue
					}

					points, err := g.countPoints(normalizedWord, len(letters), cords, horizontal)

					if err != nil {
						log.Debugf("There was error after counting points for word: %v. Error: %v", w, err)
						continue
					}
					wordsCollection = append(wordsCollection, gaddagWord{
						Cords:      cords,
						Word:       normalizedWord,
						Horizontal: horizontal,
						Points:     points},
					)

				}
				log.Debugf("There is overall %v words after counting points", len(wordsCollection))

			}
		}
	}

	return wordsCollection
}

func mockRowForStartingTile(hookIndex int, letter rune, row []rune) []rune {
	slicecopy := append([]rune(nil), row...)
	slicecopy[hookIndex] = letter
	return slicecopy
}

func prepareWordAndFixCords(word string, cords [2]int, horizontal bool) (string, [2]int) {

	i := strings.Index(word, ".")
	if i == -1 {
		return word, cords
	}
	cords[1] = cords[1] - i + 1

	// Redo cords after searching for words in transposed board
	if !horizontal {
		x := cords[0]
		cords[0] = cords[1]
		cords[1] = x
	}

	nWord := gaddag.NormalizeWord(word)

	return nWord, cords
}
