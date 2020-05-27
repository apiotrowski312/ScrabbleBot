package grabble

type field struct {
	Bonus  rune
	Letter rune
}

type Board [15][15]*field

func CreateBoard(template [15][15]rune) (Board, error) {
	var board Board

	for x, row := range template {
		for y, f := range row {
			board[x][y] = &field{Bonus: f}
		}
	}

	return board, nil
}

func (b *Board) TransposeBoard() Board {
	var transposedBoard Board

	for i, row := range b {
		for j := range row {
			transposedBoard[i][j] = b[j][i]
		}
	}
	return transposedBoard
}
