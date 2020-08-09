package board

import "github.com/apiotrowski312/scrabbleBot/utils/str_manipulator"

type field struct {
	Bonus  rune
	Letter rune
}

// TODO: Change all small letters to big one :/
type Board [15][15]*field

func CreateBoard(template [15][15]rune) *Board {

	var board Board
	for x, row := range template {
		for y, f := range row {
			board[x][y] = &field{Bonus: f}
		}
	}

	return &board
}

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
		} else if startPos[0] > 0 && b[startPos[0]-1][startPos[1]+i].Letter != rune(0) {
			hook = true
		} else if startPos[0] < 15 && b[startPos[0]+1][startPos[1]+i].Letter != rune(0) {
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

func (b *Board) placeWord(word string, startPos [2]int) {
	for i, letter := range word {
		b[startPos[0]][startPos[1]+i].Letter = letter
	}
}

// TODO: Create Bench tests
func (b *Board) GetAllWordsAndBonuses(word string, startPos [2]int, horizontal bool) ([]string, []string) {
	if horizontal {
		return b.getAllWordsAndBonuses(word, startPos)
	}

	tb := b.TransposeBoard()
	return tb.getAllWordsAndBonuses(word, [2]int{startPos[1], startPos[0]})
}

func (b *Board) getAllWordsAndBonuses(word string, startPos [2]int) ([]string, []string) {
	words := []string{word}
	bonuses := []string{""}

	for i := range word {
		if b[startPos[0]][startPos[1]+i].Letter != rune(0) {
			bonuses[0] += "0"
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
			bonus += "0"
			index++
		}

		currentWord = str_manipulator.Reverse(currentWord)
		bonus = str_manipulator.Reverse(bonus)

		index = 1
		for startPos[0]+index <= 14 && b[startPos[0]+index][startPos[1]+i].Letter != rune(0) {
			currentWord += string(b[startPos[0]+index][startPos[1]+i].Letter)
			bonus += "0"
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
