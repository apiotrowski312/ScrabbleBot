package imgPrinter

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/apiotrowski312/scrabbleBot/grabble"
)

const (
	rectSize = 50
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
	c.rect(x1, y1, x2, y2, col)

	y1++
	for ; y1 < y2; y1++ {
		c.hLine(x1+1, y1, x2-1, color.White)
	}
}

func PrintScreenBoard(g grabble.Grabble, imgName string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{1000, 1000}
	img := customImage{image.NewRGBA(image.Rectangle{upLeft, lowRight})}

	for v, row := range g.Board {
		for h := range row {
			img.fullRect(v*rectSize, h*rectSize, v*rectSize+rectSize, h*rectSize+rectSize, color.Black)
		}
	}

	f, _ := os.Create(imgName)
	png.Encode(f, img.image)
}
