package grabble

import (
	"fmt"
)

type field struct {
	Bonus  rune
	Letter rune
}

type Board [][]field

func CreateBoard(template [][]rune) (Board, error) {
	boardSizeY := len(template)
	for i, row := range template {
		if boardSizeY != len(row) {
			return nil, fmt.Errorf("Board %v row is different lenght than board hight which is %v", i, boardSizeY)
		}
	}

	var board Board

	for _, row := range template {
		var boardRow []field
		for _, f := range row {
			boardRow = append(boardRow, field{Bonus: f})
		}
		board = append(board, boardRow)
	}

	return board, nil
}

func (b Board) TransposeBoard() Board {
	var transposedBoard Board

	for i, row := range b {
		var transposedRow []field
		for j := range row {
			transposedRow = append(transposedRow, b[j][i])
		}
		transposedBoard = append(transposedBoard, transposedRow)
	}
	return transposedBoard
}
