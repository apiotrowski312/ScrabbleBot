package main

import (
	"os"

	"github.com/apiotrowski312/scrabbleBot/grabble"
	"github.com/therecipe/qt/widgets"
)

func main() {

	game := grabble.CreateDeafultGame([]string{"p1", "p2"})

	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(1000, 1000)
	window.SetWindowTitle("Grabble")

	upRow := widgets.NewQWidget(nil, 0)
	upRow.SetLayout(widgets.NewQHBoxLayout())
	upRow.Layout().AddWidget(drawBoard(game))
	upRow.Layout().AddWidget(drawPanel(game))

	downRow := widgets.NewQWidget(nil, 0)
	downRow.SetLayout(widgets.NewQHBoxLayout())

	main := widgets.NewQWidget(nil, 0)
	main.SetLayout(widgets.NewQVBoxLayout())
	main.Layout().AddWidget(upRow)
	main.Layout().AddWidget(downRow)

	window.SetCentralWidget(main)
	window.Show()
	app.Exec()
}

func drawBoard(game grabble.Grabble) widgets.QWidget_ITF {
	board := widgets.NewQWidget(nil, 0)
	board.SetLayout(widgets.NewQVBoxLayout())
	for _, row := range game.Board {
		buttonRow := widgets.NewQWidget(nil, 0)
		buttonRow.SetLayout(widgets.NewQHBoxLayout())
		for _, cell := range row {
			button := widgets.NewQPushButton2(string(cell.Letter), nil)
			button.SetFixedWidth(50)
			button.SetFixedHeight(50)
			button.SetStyleSheet("background-image:url(\"frontend/icons/" + string(cell.Bonus) + ".jpg\"); background-position: center;")
			buttonRow.Layout().AddWidget(button)
		}
		board.Layout().AddWidget(buttonRow)
	}
	return board
}

func drawPanel(game grabble.Grabble) widgets.QWidget_ITF {
	panel := widgets.NewQWidget(nil, 0)
	panel.SetLayout(widgets.NewQVBoxLayout())
	title := widgets.NewQLabel2("Points:", nil, 1)
	panel.Layout().AddWidget(title)
	for _, player := range game.Players {
		label := widgets.NewQLabel2(player.Name+": "+string(player.Points)+"_", nil, 1)
		panel.Layout().AddWidget(label)
	}
	return panel
}
