package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type timerPage struct {
	app  *application
	grid *tview.Grid
}

var (
	timerValueInMins int = 30
)

var (
	workSessionStage int
)

const (
	notStarted int = iota
	inWork
	onPause
	ended
)

//	_____________________________
//	|__________title____________|
//	|			|				|
//	|			|				|
//	|	info	|	timer		|
//	|			|				|
//	|___________|_______________|
//	|__________footer___________|

func newTimerPage(app *application) *timerPage {
	page := tview.NewGrid()
	page.SetColumns(0)
	page.SetRows(1, 0, 1)

	title := tview.NewTextView()
	title.SetText("TIMER PAGE")
	title.SetTextAlign(tview.AlignCenter)

	layout := tview.NewGrid()
	layout.SetColumns(0, 0)
	layout.SetRows(0)
	layout.SetBorder(true)

	layoutInfo := tview.NewTextView()
	layoutInfo.SetTextAlign(tview.AlignCenter)

	layoutTimer := tview.NewTextView()
	layoutTimer.SetTextAlign(tview.AlignCenter)
	layoutTimer.SetScrollable(false)

	layout.AddItem(layoutInfo, 0, 0, 1, 1, 0, 0, false)
	layout.AddItem(layoutTimer, 0, 1, 1, 1, 0, 0, false)

	footer := tview.NewTextView()
	footer.SetText("Press [e] for end ")
	footer.SetTextAlign(tview.AlignCenter)

	page.AddItem(title, 0, 0, 1, 1, 0, 0, false)
	page.AddItem(layout, 1, 0, 1, 1, 0, 0, true)
	page.AddItem(footer, 2, 0, 1, 1, 0, 0, false)

	exitChan := make(chan struct{}, 1)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'e':
			if workSessionStage == inWork || workSessionStage == onPause {
				exitChan <- struct{}{}
			}
			app.switchToMainPage()
		}
		return event
	})

	layout.SetFocusFunc(func() {

		switch workSessionStage {
		case onPause, inWork:
			return
		case notStarted, ended:
			break
		default:
			return
		}

		workSessionStage = inWork

		// TODO: Change to minutes
		timerDuration := time.Minute * time.Duration(timerValueInMins)
		timer := time.NewTimer(timerDuration)
		ticker := time.NewTicker(time.Second)

		startTime := time.Now()
		endTime := startTime.Add(timerDuration)
		layoutInfo.SetText(
			fmt.Sprintf(
				"Start work session at: %s\nSession ended in: %s",
				startTime.Format(time.TimeOnly),
				endTime.Format(time.TimeOnly),
			),
		)

		go func() {
			var stopped bool
			for !stopped {
				select {
				case <-exitChan:
					timer.Stop()
					ticker.Stop()

					workSessionStage = ended

					app.QueueUpdateDraw(func() {
						layoutTimer.Clear()
						layoutInfo.Clear()
					})

					stopped = true
				case t := <-ticker.C:
					app.QueueUpdateDraw(func() {
						tt := time.Time{}
						tt = tt.Add(endTime.Sub(t))

						layoutTimer.SetText(tt.Format(time.TimeOnly))
					})
				case t := <-timer.C:
					ticker.Stop()

					workSessionStage = ended

					

					app.QueueUpdateDraw(func() {
						layoutTimer.SetText("Game Over at " + t.Format(time.TimeOnly))
					})
					stopped = true
				}
			}
		}()
	})

	return &timerPage{
		app:  app,
		grid: page,
	}
}
