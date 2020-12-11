package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func loadUI() fyne.CanvasObject {
	list := widget.NewVBox(
		widget.NewLabel("Item 1"),
		widget.NewLabel("Item 2"),
	)

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	content := widget.NewMultiLineEntry()

	left := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, list)

	ui := widget.NewHSplitContainer(left, content)
	ui.Offset = 0.25
	return ui
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	w.SetContent(loadUI())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
