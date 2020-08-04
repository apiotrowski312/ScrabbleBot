package grabble

// TODO: Function with starting word

type wordCollection struct {
	cords      [2]int
	words      []string
	horizontal bool
}

type bestWord struct {
	Points     int
	Cords      [2]int
	Word       string
	Horizontal bool
}

func (g *Grabble) PickBestWord() bestWord {

	rack := g.CurrentPlayer().Rack

	wordsCollection := g.getWordCollection(rack, true)
	wordsCollection = append(wordsCollection, g.getWordCollection(rack, false)...)

	bestW := bestWord{}
	for _, singleCollection := range wordsCollection {
		for _, word := range singleCollection.words {
			// BUG: Its not working properly. I should pass not a rack, but used letters.
			if points, err := g.countPoints(word, singleCollection.cords, true); err == nil && points > bestW.Points {
				bestW = bestWord{
					Cords:  singleCollection.cords,
					Word:   word,
					Points: points,
				}
			}
		}
	}

	return bestW
}

func (g *Grabble) getWordCollection(rack []rune, horizontal bool) []wordCollection {
	board := g.Board
	if !horizontal {
		board = *g.Board.TransposeBoard()
	}

	wordsCollection := []wordCollection{}

	for x, row := range board {
		rowLetters := g.getRowOfLetters(x)
		for y, _ := range row {
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
				wordsCollection = append(wordsCollection, wordCollection{
					cords:      [2]int{x, y},
					words:      words,
					horizontal: horizontal})
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
