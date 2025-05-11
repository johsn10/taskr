package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/johsn10/taskr/gui/editor"
)

var app *tview.Application
var pages *tview.Pages
var activePage string

func Init() {
	app = tview.NewApplication()
	pages = tview.NewPages()
	redraw := func() {
		app.Draw()
	}

	pages.AddPage("editor", editor.UI(redraw), true, true)
	activePage = "editor"

	app.SetInputCapture(eventHandler)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func switchToPage(page string) {
	activePage = page
	pages.SwitchToPage(page)
}

func eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEscape {
		app.Stop()
		return nil
	}

	if activePage == "editor" {
		return editor.HandleEvents(eventKey)
	}

	return eventKey
}
