package gui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func initUI(app fyne.App) {
	//var login structs.LoginT

	w := app.NewWindow("Login")

	content := fyne.NewContainer(
		canvas.NewText("Login:", color.RGBA{0, 0x80, 0, 0xff}))
	content.Layout = layout.NewFixedGridLayout(fyne.NewSize(400, 400))

	w.SetContent(content)
	w.Show()
}

// Show draws login view
func Show(app fyne.App) {
	initUI(app)
}
