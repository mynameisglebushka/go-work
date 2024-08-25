package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mainPage struct {
	app  *application
	grid *tview.Grid
}

func newMainPage(app *application) *mainPage {
	page := tview.NewGrid()

	title := tview.NewTextView()
	title.SetText("MAIN PAGE")
	title.SetTextAlign(tview.AlignCenter)

	layout := tview.NewForm()
	layout.SetTitle("Set Work Session Settings")
	layout.AddInputField("Work Session lenght in minuts", "30", 0, func(textToCheck string, lastChar rune) bool {
		minuts, err := strconv.Atoi(textToCheck)
		if err != nil {
			return false
		}

		if minuts <= 0 {
			return false
		}

		return true
	}, func(text string) {
		v, err := strconv.Atoi(text)
		if err != nil {
			return
		}

		if v <= 0 {
			return
		}

		timerValueInMins = v
	})
	layout.AddButton("Start", func() {
		app.switchToTimerPage()
	})

	footer := tview.NewTextView()
	footer.SetText("Press [q] for quit")
	footer.SetTextAlign(tview.AlignCenter)

	page.AddItem(title, 0, 0, 1, 1, 0, 0, false)
	page.AddItem(layout, 1, 0, 1, 1, 0, 0, true)
	page.AddItem(footer, 2, 0, 1, 1, 0, 0, false)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
		}
		return event
	})

	return &mainPage{
		app:  app,
		grid: page,
	}
}
