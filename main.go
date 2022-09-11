package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	editWidget := widget.NewMultiLineEntry()
	previewWidget := widget.NewRichTextFromMarkdown("")

	app.EditWidget = editWidget
	app.PreviewWidget = previewWidget

	editWidget.OnChanged = previewWidget.ParseMarkdown

	return editWidget, previewWidget
}

func main() {
	// create a fyne app
	fyneApp := app.New()

	// create a window for the app
	win := fyneApp.NewWindow("Markdown Editor")

	// get the user interface
	editWidget, previewWidget := cfg.makeUI()

	// set the content of the window
	win.SetContent(container.NewHSplit(editWidget, previewWidget))

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()

}
