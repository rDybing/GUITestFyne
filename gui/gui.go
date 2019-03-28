package gui

import (

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func initUI(app fyne.App) {
	//var login structs.LoginT

	w := app.NewWindow("Login")
	w.SetContent(widget.NewHBox(makeMainMenu(app)))
	content.Layout = layout.NewFixedGridLayout(fyne.NewSize(400, 400))

	w.Show()
}

func makeMainMenu(app fyne.App) *fyne.Container {
	users := widget.NewButton("Users", func() {
		viewName := widget.NewLabel("Users")
		horisontal := widget.NewHBox(makeMainMenu(), makeUsersMenu(app), viewName)
	})
	networks := widget.NewButton("Networks", func() {
		viewName := widget.NewLabel("Networks")
		horisontal := widget.NewHBox(makeMainMenu(), makeNetworksMenu(app), viewName)
	})
	groups := widget.NewButton("Groups", func() {
		app.Quit()
	})
	auxApps := widget.NewButton("AuxApps", func() {
		app.Quit()
	})
	logs := widget.NewButton("Log", func() {
		app.Quit()
	})
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), users, networks, groups, auxApps, logs)
}

func makeUsersMenu(app fyne.App) *fyne.Container {
	recent := widget.NewButton("Recent", func() {
		App.Quit()
	})
	search := widget.NewButton("Search", func() {
		App.Quit()
	})
	back := widget.NewButton("Back", func() {
		App.Quit()
	})
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), recent, search, back)
}

func makeNetworksMenu() *fyne.Container {
	all := widget.NewButton("All", func() {
		App.Quit()
	})
	prem := widget.NewButton("Premium", func() {
		App.Quit()
	})
	premPlus := widget.NewButton("Premium+", func() {
		App.Quit()
	})
	free := widget.NewButton("Free", func() {
		App.Quit()
	})
	back := widget.NewButton("Back", func() {
		App.Quit()
	})
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), all, prem, premPlus, free, back)
}

// Show draws login view
func Show(app fyne.App) {
	initUI(app)
}
