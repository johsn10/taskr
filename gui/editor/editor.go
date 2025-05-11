package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func UI(redraw func()) *tview.Grid {

	textArea := tview.NewTextArea()
	textArea.SetBorder(true)
	textArea.SetTitle("TaskR")
	textArea.SetBorderPadding(1, 1, 3, 3)
	textArea.SetPlaceholder("Visit github.com/johsn10/taskr to learn more about taskr-script")
	textArea.SetPlaceholderStyle(
		textArea.GetPlaceholderStyle().
			Url("https://github.com/johsn10/taskr").
			Italic(true).Foreground(tcell.ColorDarkGray),
	)

	helpInfo := tview.NewTextView()
	helpInfo.SetText("Press Esc to quit, press Ctr+r to run, press Ctr+h to hault, and press Ctr+s to save")

	grid := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false)

	return grid
}

func HandleEvents(eventKey *tcell.EventKey) *tcell.EventKey {
	return eventKey
}
