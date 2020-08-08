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
	rack := g.CurrentPlayer().Rack

	wordsCollection := g.getWordCollection(rack, true)
	wordsCollection = append(wordsCollection, g.getWordCollection(rack, false)...)

	sort.Slice(wordsCollection, func(i, j int) bool {
		return wordsCollection[i].Points > wordsCollection[j].Points
	})

	if len(wordsCollection) < numberOfWords {
		return wordsCollection
	}
	return wordsCollection[:numberOfWords]
}

func (g *Grabble) getWordCollection(rack []rune, horizontal bool) []gaddagWord {
	board := g.Board
	if !horizontal {
		board = *g.Board.TransposeBoard()
	}

	wordsCollection := []gaddagWord{}

	for x, row := range board {
		rowLetters := g.getRowOfLetters(x)
		for y := range row {
			words := []string{}
			if board[x][y].Letter != rune(0) && y > 0 && board[x][y-1].Letter == rune(0) {
				words = g.Dict.FindAllWords(y, rowLetters, rack)
			} else if board[x][y].Bonus == 's' {
				// HACK: Better and faster option will be create new function in gaddag to look for words without hook
				for _, l := range rack {
					sWords := g.Dict.FindAllWords(y, mockRowWithHookWhenStartingLetter(y, l, rowLetters), rack)

					words = append(words, sWords...)
				}
			}
			if len(words) != 0 {
				for _, w := range words {
					normalizedWord, cords := prepareWord(w, [2]int{x, y}, horizontal)

					letters, letterError := g.validateAndExtractUsedNewLetters(normalizedWord, cords, horizontal)
					if letterError != nil {
						continue
					}

					if points, err := g.countPoints(normalizedWord, len(letters), cords, horizontal); err == nil {
						wordsCollection = append(wordsCollection, gaddagWord{
							Cords:      cords,
							Word:       normalizedWord,
							Horizontal: horizontal,
							Points:     points},
						)
					}

				}
			}
		}
	}

	return wordsCollection
}

func mockRowWithHookWhenStartingLetter(hookIndex int, letter rune, row []rune) []rune {
	slicecopy := append([]rune(nil), row...)
	slicecopy[hookIndex] = letter
	return slicecopy
}

func (g *Grabble) getRowOfLetters(row int) []rune {
	letters := []rune{}
	for _, letter := range g.Board[row] {
		letters = append(letters, letter.Letter)
	}
	return letters
}

func prepareWord(word string, cords [2]int, horizontal bool) (string, [2]int) {
	i := strings.Index(word, ".")
	if i == -1 {
		return word, cords
	}

	if horizontal {
		cords[1] = cords[1] - i + 1
	} else {
		cords[0] = cords[0] - i + 1
	}
	nWord := gaddag.NormalizeWord(word)

	return nWord, cords
}
