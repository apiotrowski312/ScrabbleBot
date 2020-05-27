package grabble

type field struct {
	Bonus  rune
	Letter rune
}

type Board [15][15]*field

var (
	board           = Board{}
	transposedBoard = Board{}
)

func CreateBoard(template [15][15]rune) *Board {
	if (Board{}) != board {
		return &board
	}

	for x, row := range template {
		for y, f := range row {
			board[x][y] = &field{Bonus: f}
		}
	}

	return &board
}

func (b *Board) TransposeBoard() *Board {
	if (Board{}) != transposedBoard {
		return &transposedBoard
	}

	for i, row := range b {
		for j := range row {
			transposedBoard[i][j] = b[j][i]
		}
	}
	return &transposedBoard
}
