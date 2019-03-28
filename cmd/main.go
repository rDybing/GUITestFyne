package main

import (
	"fyne.io/fyne/app"
	"github.com/rDybing/GUITestFyne/gui"
)

func main() {
	app := app.New()
	gui.Show(app)
	app.Run()
}
