package img_printer

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"unicode"

	"github.com/apiotrowski312/scrabbleBot/grabble"
)

var (
	rectSize   = 15
	screenSize = [2]int{350, 226}
	colors     = map[string]color.RGBA{
		"letter": {0, 0, 0, 255},
		"tile":   {255, 255, 255, 255},
		"w":      {204, 255, 204, 255},
		"W":      {102, 255, 102, 255},
		"l":      {204, 204, 255, 255},
		"L":      {102, 102, 255, 255},
		"s":      {204, 255, 204, 255},
		"":       {0, 0, 0, 255},
	}
)

type customImage struct {
	image *image.RGBA
}

func (c customImage) hLine(x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		c.image.Set(x1, y, col)
	}
}

func (c customImage) vLine(x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		c.image.Set(x, y1, col)
	}
}

func (c customImage) rect(x1, y1, x2, y2 int, col color.Color) {
	c.hLine(x1, y1, x2, col)
	c.hLine(x1, y2, x2, col)
	c.vLine(x1, y1, y2, col)
	c.vLine(x2, y1, y2, col)
}

func (c customImage) fullRect(x1, y1, x2, y2 int, col color.Color) {
	for ; y1 < y2; y1++ {
		c.hLine(x1, y1, x2, col)
	}
}

func (c customImage) drawTileFromArray(x, y int, tile [][]int) {
	for h, row := range tile {
		for v, cell := range row {
			if cell == 1 {
				c.image.Set(x+v, y+h, colors["letter"])
			}
		}
	}
}

func (c customImage) drawBonus(x, y int, tileType rune) {
	switch tileType {
	case 'l':
		c.drawTileFromArray(x, y, bonus2xLetter)
	case 'L':
		c.drawTileFromArray(x, y, bonus3xLetter)
	case 'w':
		c.drawTileFromArray(x, y, bonus2xWord)
	case 'W':
		c.drawTileFromArray(x, y, bonus3xWord)
	case 's':
		c.drawTileFromArray(x, y, startTile)
	}
}

func (c customImage) drawLetter(x, y int, letter rune) {
	switch letter {
	case 'A':
		c.drawTileFromArray(x, y, letterA)
	case 'B':
		c.drawTileFromArray(x, y, letterB)
	case 'C':
		c.drawTileFromArray(x, y, letterC)
	case 'D':
		c.drawTileFromArray(x, y, letterD)
	case 'E':
		c.drawTileFromArray(x, y, letterE)
	case 'F':
		c.drawTileFromArray(x, y, letterF)
	case 'G':
		c.drawTileFromArray(x, y, letterG)
	case 'H':
		c.drawTileFromArray(x, y, letterH)
	case 'I':
		c.drawTileFromArray(x, y, letterI)
	case 'J':
		c.drawTileFromArray(x, y, letterJ)
	case 'K':
		c.drawTileFromArray(x, y, letterK)
	case 'L':
		c.drawTileFromArray(x, y, letterL)
	case 'M':
		c.drawTileFromArray(x, y, letterM)
	case 'N':
		c.drawTileFromArray(x, y, letterN)
	case 'O':
		c.drawTileFromArray(x, y, letterO)
	case 'P':
		c.drawTileFromArray(x, y, letterP)
	case 'Q':
		c.drawTileFromArray(x, y, letterQ)
	case 'R':
		c.drawTileFromArray(x, y, letterR)
	case 'S':
		c.drawTileFromArray(x, y, letterS)
	case 'T':
		c.drawTileFromArray(x, y, letterT)
	case 'U':
		c.drawTileFromArray(x, y, letterU)
	case 'V':
		c.drawTileFromArray(x, y, letterV)
	case 'W':
		c.drawTileFromArray(x, y, letterW)
	case 'X':
		c.drawTileFromArray(x, y, letterX)
	case 'Y':
		c.drawTileFromArray(x, y, letterY)
	case '_':
		c.drawTileFromArray(x, y, letter_)
	case 'Z':
		c.drawTileFromArray(x, y, letterZ)
	case '0':
		c.drawTileFromArray(x, y, number0)
	case '1':
		c.drawTileFromArray(x, y, number1)
	case '2':
		c.drawTileFromArray(x, y, number2)
	case '3':
		c.drawTileFromArray(x, y, number3)
	case '4':
		c.drawTileFromArray(x, y, number4)
	case '5':
		c.drawTileFromArray(x, y, number5)
	case '6':
		c.drawTileFromArray(x, y, number6)
	case '7':
		c.drawTileFromArray(x, y, number7)
	case '8':
		c.drawTileFromArray(x, y, number8)
	case '9':
		c.drawTileFromArray(x, y, number9)
	}
}

func PrintScreenBoard(g grabble.Grabble, imgName string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{screenSize[0], screenSize[1]}
	img := customImage{image.NewRGBA(image.Rectangle{upLeft, lowRight})}

	img.fullRect(0, 0, screenSize[0], screenSize[1], colors["tile"])

	// Draw board
	for y, row := range g.Board {
		for x, cell := range row {
			img.rect(x*rectSize, y*rectSize, x*rectSize+rectSize, y*rectSize+rectSize, colors["letter"])
			if cell.Bonus != rune(0) {
				img.fullRect(x*rectSize+1, y*rectSize+1, x*rectSize+rectSize-1, y*rectSize+rectSize, colors[string(cell.Bonus)])
			}
			img.drawBonus(x*rectSize+5, y*rectSize+9, cell.Bonus)
			img.drawLetter(x*rectSize+2, y*rectSize+2, unicode.ToUpper(cell.Letter))
		}
	}

	xAxis := 230
	// Draw round:
	for x, letter := range "Round:" {
		img.drawLetter(xAxis+x*6, 1, unicode.ToUpper(letter))
	}
	for x, letter := range strconv.Itoa(g.Stats.CurrentRound) {
		img.drawLetter(xAxis+50+x*6, 1, unicode.ToUpper(letter))
	}
	// Draw racks
	yRack := 10
	for x, letter := range "Racks:" {
		img.drawLetter(xAxis+x*6, yRack+5, unicode.ToUpper(letter))
	}
	for y, player := range g.Players {
		for x, letter := range player.Name {
			img.drawLetter(xAxis+x*6, yRack+10+y*10+5, unicode.ToUpper(letter))
		}
		for x, letter := range player.Rack {
			img.drawLetter(xAxis+50+x*6, yRack+10+y*10+5, unicode.ToUpper(letter))
		}
	}

	// Draw scores
	yScores := 100
	for x, letter := range "Score:" {
		img.drawLetter(xAxis+x*6, yScores+5, unicode.ToUpper(letter))
	}

	for y, player := range g.Players {
		for x, letter := range player.Name {
			img.drawLetter(xAxis+x*6, yScores+10+y*10+5, unicode.ToUpper(letter))
		}
		for x, letter := range strconv.Itoa(player.Points) {
			img.drawLetter(xAxis+50+x*6, yScores+10+y*10+5, unicode.ToUpper(letter))
		}
	}

	// Draw winner
	yWinner := 180
	if g.Stats.Finished {
		for x, letter := range "Winner:" {
			img.drawLetter(xAxis+x*6, yWinner+5, unicode.ToUpper(letter))
		}
		for x, letter := range g.Stats.Winner.Name {
			img.drawLetter(xAxis+50+x*6, yWinner+5, unicode.ToUpper(letter))
		}
		for x, letter := range strconv.Itoa(g.Stats.Winner.Points) {
			img.drawLetter(xAxis+100+x*6, yWinner+5, unicode.ToUpper(letter))
		}
	}

	f, _ := os.Create(imgName)
	png.Encode(f, img.image)
}
