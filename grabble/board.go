package grabble

type field struct {
	Bonus  rune
	Letter rune
}

type Board [15][15]*field

var ()

func CreateBoard(template [15][15]rune) *Board {

	var board Board
	for x, row := range template {
		for y, f := range row {
			board[x][y] = &field{Bonus: f}
		}
	}

	return &board
}

// TODO: Create little helper package to pretty print fixtures and golden files

func (b *Board) TransposeBoard() *Board {
	var transposedBoard Board
	for i, row := range b {
		for j := range row {
			transposedBoard[i][j] = b[j][i]
		}
	}
	return &transposedBoard
}

func (b *Board) CanWordBePlaced(word string, startPos [2]int, horizontal bool) bool {
	if horizontal {
		return b.canWordBePlaced(word, startPos)
	}

	tb := b.TransposeBoard()
	return tb.canWordBePlaced(word, [2]int{startPos[1], startPos[0]})
}

func (b *Board) canWordBePlaced(word string, startPos [2]int) bool {
	hook := false

	if startPos[1]+len(word) > 15 {
		return false
	}

	for i, letter := range word {
		bc := b[startPos[0]][startPos[1]+i]

		if bc.Letter != rune(0) && bc.Letter != letter {
			return false
		}

		if bc.Bonus == rune('s') || bc.Letter == letter {
			hook = true
		} else if startPos[0] > 0 && b[startPos[0]-1][startPos[1]+i].Letter != rune(0) {
			hook = true
		} else if startPos[0] < 15 && b[startPos[0]+1][startPos[1]+i].Letter != rune(0) {
			hook = true
		}
	}

	if hook == false {
		return false
	}

	return true
}

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

// TODO: Create helper function for getting board/transposedBoard
// TODO: Collect all new words
