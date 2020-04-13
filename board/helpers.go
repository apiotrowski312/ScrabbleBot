package board

func (b Board) getTileType(cord [2]int) rune {
	if b[cord[0]][cord[1]].Letter != rune(0) {
		return '0'
	}
	return b[cord[0]][cord[1]].TileType
}
