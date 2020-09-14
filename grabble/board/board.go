package board

import "github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"

type field struct {
	Bonus  rune
	Letter rune
}

// Board struct - simple 15x15 matrix representing state of board
type Board [15][15]*field

// CreateBoard - create board from a 15x15 template.
// Temple use following pattern:
// * 0 - just normal, regular tile
// * s - starting tile, there can be more than one of them
// * l - double letter bonus
// * L - triple letter bonus
// * w - double word bonus
// * W - triple word bonus
func CreateBoard(template [15][15]rune) *Board {
	var board Board
	for x, row := range template {
		for y, f := range row {
			board[x][y] = &field{Bonus: f}
		}
	}

	return &board
}

// TransposeBoard - transpose board, makes other functions easier to write
// IMPORTANT! In transposed board cords are reversed
func (b *Board) TransposeBoard() *Board {
	var transposedBoard Board
	for i, row := range b {
		for j := range row {
			transposedBoard[i][j] = b[j][i]
		}
	}
	return &transposedBoard
}

// DoesHookExist - return array with all new letters to place
// and boolean with information if word can be placed on the board.
// False is e.g. when in conflict with existing letters on board
func (b *Board) DoesHookExist(word string, startPos [2]int, horizontal bool) ([]rune, bool) {
	if horizontal {
		return b.doesHookExist(word, startPos)
	}

	tb := b.TransposeBoard()
	return tb.doesHookExist(word, [2]int{startPos[1], startPos[0]})
}

func (b *Board) doesHookExist(word string, startPos [2]int) ([]rune, bool) {
	hook := false

	// FIXME: Avoid hardcodeing numbers like 15. What number is that?
	if startPos[1]+len(word) >= 15 || startPos[1] < 0 {
		return []rune{}, false
	}

	if startPos[1] > 0 && b[startPos[0]][startPos[1]-1].Letter != rune(0) {
		return []rune{}, false
	}

	if startPos[1]+len(word) < 14 && b[startPos[0]][startPos[1]+len(word)+1].Letter != rune(0) {
		return []rune{}, false
	}

	newLetters := []rune{}

	for i, letter := range word {
		bc := b[startPos[0]][startPos[1]+i]

		if bc.Letter != rune(0) && bc.Letter != letter {
			return []rune{}, false
		}
		if bc.Letter == letter {
			hook = true
			continue
		}

		newLetters = append(newLetters, letter)
		if bc.Bonus == rune('s') {
			hook = true
		} else if startPos[0] >= 1 && b[startPos[0]-1][startPos[1]+i].Letter != rune(0) {
			hook = true
		} else if startPos[0] <= 13 && b[startPos[0]+1][startPos[1]+i].Letter != rune(0) {
			hook = true
		}
	}

	if hook == false {
		return []rune{}, false
	}

	return newLetters, true
}

// PlaceWord - will place word on board. Function assumes that there is no conflicts on board.
// If there will be any, it will overwrite existing letters.
func (b *Board) PlaceWord(word string, startPos [2]int, horizontal bool) {
	if horizontal {
		b.placeWord(word, startPos)
	} else {
		tb := b.TransposeBoard()
		tb.placeWord(word, [2]int{startPos[1], startPos[0]})
	}
}

// TODO: Blank - placeWord need new flag (if blank and where)
// Do not place letter if it is already places
// Add information aboyt round and who put it
func (b *Board) placeWord(word string, startPos [2]int) {
	for i, letter := range word {
		b[startPos[0]][startPos[1]+i].Letter = letter
	}
}

// GetAllWordsAndBonuses gets all new words(with bonuses) created with new placed word
// TODO: Create Bench tests
func (b *Board) GetAllWordsAndBonuses(word string, startPos [2]int, horizontal bool) ([]string, []string) {
	if horizontal {
		return b.getAllWordsAndBonuses(word, startPos)
	}

	tb := b.TransposeBoard()
	return tb.getAllWordsAndBonuses(word, [2]int{startPos[1], startPos[0]})
}

// TODO: Blank - blank should have 0 points.
// Maybe add bool variable to field to mark if blank was used here?
func (b *Board) getAllWordsAndBonuses(word string, startPos [2]int) ([]string, []string) {
	words := []string{word}
	bonuses := []string{""}

	for i := range word {
		if b[startPos[0]][startPos[1]+i].Letter != rune(0) {
			bonuses[0] += string(rune(0))
		} else {
			bonuses[0] += string(b[startPos[0]][startPos[1]+i].Bonus)
		}
	}

	for i := range word {
		if b[startPos[0]][startPos[1]+i].Letter != rune(0) {
			continue
		}

		currentWord := string(word[i])
		bonus := string(b[startPos[0]][startPos[1]+i].Bonus)

		index := 1
		for startPos[0]-index >= 0 && b[startPos[0]-index][startPos[1]+i].Letter != rune(0) {
			currentWord += string(b[startPos[0]-index][startPos[1]+i].Letter)
			bonus += string(rune(0))
			index++
		}

		currentWord = str_manipulator.Reverse(currentWord)
		bonus = str_manipulator.Reverse(bonus)

		index = 1
		for startPos[0]+index <= 14 && b[startPos[0]+index][startPos[1]+i].Letter != rune(0) {
			currentWord += string(b[startPos[0]+index][startPos[1]+i].Letter)
			bonus += string(rune(0))
			index++
		}

		// len == 1 mean there was no word in this column
		if len(currentWord) == 1 {
			continue
		}

		words = append(words, currentWord)
		bonuses = append(bonuses, bonus)
	}

	return words, bonuses
}

// GetRowOfLetters returns row of letters as a rune array.
func (b *Board) GetRowOfLetters(row int) []rune {
	letters := []rune{}
	for _, letter := range b[row] {
		letters = append(letters, letter.Letter)
	}
	return letters
}
