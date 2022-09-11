package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

func (app *config) createMenuItems(win fyne.Window) {

	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func main() {
	// create a fyne app
	fyneApp := app.New()

	// create a window for the app
	win := fyneApp.NewWindow("Markdown Editor")

	// get the user interface
	editWidget, previewWidget := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(editWidget, previewWidget))

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()

}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
			}

			if write == nil {
				// user canceled
				return
			}

			// save the file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.Show()
	}
}
