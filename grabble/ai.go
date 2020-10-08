package grabble

import (
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/apiotrowski312/scrabbleBot/gaddag"
	"github.com/apiotrowski312/scrabbleBot/grabble/board"
)

// TODO: Function with starting word
// FIXME: Better naming for this variable
type gaddagWord struct {
	Points     int
	Cords      [2]int
	Word       string
	Horizontal bool
	Ratio      float64
}

// PickBestWord returns best words for current player and board
func (g *Grabble) PickBestWord(numberOfWords int) []gaddagWord {
	log.Debugf("PickBestWord function called by %s", g.CurrentPlayer().Name)
	log.Debugf("Rack: %s", string(g.CurrentPlayer().Rack))
	rack := g.CurrentPlayer().Rack

	wordChan := make(chan []gaddagWord, 2)
	var wg sync.WaitGroup
	for _, h := range [2]bool{true, false} {
		wg.Add(1)
		go func(horizontal bool) {
			defer wg.Done()
			wc := g.getWordCollection(rack, horizontal)
			if len(wc) > 0 {
				wordChan <- wc
			}
		}(h)
	}

	wg.Wait()
	close(wordChan)

	wordsCollection := []gaddagWord{}
	for w := range wordChan {
		wordsCollection = append(wordsCollection, w...)
	}

	log.Debugf("Found %v words", len(wordsCollection))
	sort.Slice(wordsCollection, func(i, j int) bool {
		return wordsCollection[i].Ratio > wordsCollection[j].Ratio
	})

	if len(wordsCollection) < numberOfWords {
		return wordsCollection
	}
	return wordsCollection[:numberOfWords]
}

func (g Grabble) getWordCollection(rack []rune, horizontal bool) []gaddagWord {
	log.Debugf("getWordCollection called. Horizontal: %v", horizontal)

	board := &g.Board
	if !horizontal {
		board = g.Board.TransposeBoard()
	}

	wordsCollection := []gaddagWord{}

	wordChan := make(chan []gaddagWord, len(board))
	var wg sync.WaitGroup

	for x := range board {
		wg.Add(1)
		go func(x int) {
			wc := []gaddagWord{}

			for y := range board[x] {
				wc = append(wc, g.getWordsFromATile(x, y, board, horizontal)...)
			}

			if len(wc) > 0 {
				wordChan <- wc

			}
			wg.Done()
		}(x)
	}
	wg.Wait()

	close(wordChan)

	for w := range wordChan {
		wordsCollection = append(wordsCollection, w...)
	}

	return wordsCollection
}

func (g Grabble) getWordsFromATile(x, y int, board *board.Board, horizontal bool) []gaddagWord {
	words := g.getWordFromAHook(x, y, board, horizontal)
	var wordsCollection []gaddagWord
	if len(words) > 0 {
		wordsCollection = g.checkWord(words, x, y, horizontal)
	}
	return wordsCollection
}

func hookType1And2(x, y int, board *board.Board) bool {
	return board[x][y].Letter != rune(0) && (y == 0 || (y > 0 && board[x][y-1].Letter == rune(0)))
}

func hookType3(x, y int, board *board.Board) bool {
	return board[x][y].Letter == rune(0) && ((x > 0 && board[x-1][y].Letter != rune(0)) || (x < 14 && board[x+1][y].Letter != rune(0)))
}

func hookStarterTile(x, y int, board *board.Board) bool {
	return board[x][y].Bonus == 's' && board[x][y].Letter == rune(0)
}

func (g Grabble) getWordFromAHook(x, y int, board *board.Board, horizontal bool) []string {
	rack := g.CurrentPlayer().Rack
	rowLetters := board.GetRowOfLetters(x)
	words := []string{}
	if hookType1And2(x, y, board) {
		log.Debugf("Hook (type 1/2) found %v(%v). Horizontal: %v", string(board[x][y].Letter), [2]int{x, y}, horizontal)
		log.Debugf("Row for finding words %v, rack %v, hookIndex %v", rowLetters, rack, y)
		words = g.Dict.FindAllWords(y, rowLetters, rack)
	} else if hookType3(x, y, board) {
		log.Debugf("Hook (type 3) found %v(%v). Horizontal: %v", string(board[x][y].Letter), [2]int{x, y}, horizontal)
		log.Debugf("Row for finding words %v, rack %v, hookIndex %v", rowLetters, rack, y)
		// FIXME: Create proper solution for this case. There is no point of looking for a word, if we wont be able to do a word with
		// Lower/higher already exisitng word. Follwoing sollution is good enough only for now
		for i, l := range rack {
			rackForThisItteration := append(append([]rune{}, rack[:i]...), rack[i+1:]...)
			sWords := g.Dict.FindAllWords(y, mockRowForStartingTile(y, l, rowLetters), rackForThisItteration)

			words = append(words, sWords...)
		}
	} else if hookStarterTile(x, y, board) {
		log.Debugf("Starting tile found %v(%v). Horizontal: %v", string(board[x][y].Letter), [2]int{x, y}, horizontal)
		log.Debugf("Row for finding words %v, rack %v, hookIndex %v", rowLetters, rack, y)
		for i, l := range rack {
			rackForThisItteration := append(append([]rune{}, rack[:i]...), rack[i+1:]...)
			sWords := g.Dict.FindAllWords(y, mockRowForStartingTile(y, l, rowLetters), rackForThisItteration)

			words = append(words, sWords...)
		}
	}
	return words
}

func (g Grabble) checkWord(words []string, x, y int, horizontal bool) []gaddagWord {
	log.Debugf("There is %v new words before counting points", len(words))
	wordsCollection := []gaddagWord{}
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

		// MAYBE: Add constructor for gaddagWord?
		wordsCollection = append(wordsCollection, gaddagWord{
			Cords:      cords,
			Word:       normalizedWord,
			Horizontal: horizontal,
			Points:     points,
			Ratio:      getRatio(letters, points),
		},
		)

	}
	log.Debugf("There is overall %v words after counting points", len(wordsCollection))
	return wordsCollection
}

func getRatio(letters []rune, points int) float64 {
	ratio := 0.0

	for _, l := range letters {
		if unicode.IsLower(l) {
			ratio += letterRatio['_']
		} else {
			ratio += letterRatio[l]
		}
	}

	ratio /= float64(len(letters))

	return ratio * float64(points)
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
