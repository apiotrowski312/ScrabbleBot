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

// TODO: Collect all new words
// TODO: If word can be placed
