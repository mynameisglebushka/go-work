package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	mainPageName  = "main"
	timerPageName = "timer"
)

type application struct {
	*tview.Application
	pages *tview.Pages

	mainPage  *mainPage
	timerPage *timerPage
}

func newApplication() *application {
	return &application{
		Application: tview.NewApplication(),
		pages: tview.NewPages(),
	}
}

var nilMouseCapture = func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	return nil, -1
}

func (a *application) run() error {
	return a.SetRoot(a.pages, true).EnableMouse(false).SetMouseCapture(nilMouseCapture).Run()
}

func (a *application) addMainPage(page *mainPage) *application {
	a.pages.AddPage(mainPageName, page.grid, true, true)
	a.mainPage = page

	return a
}

func (a *application) addTimerPage(page *timerPage) *application {
	a.pages.AddPage(timerPageName, page.grid, true, false)
	a.timerPage = page
	return a
}

func (a *application) switchToMainPage() *application {
	a.pages.SwitchToPage(mainPageName)
	a.SetFocus(a.mainPage.grid)
	return a
}

func (a *application) switchToTimerPage() *application {
	a.pages.SwitchToPage(timerPageName)
	a.SetFocus(a.timerPage.grid)
	return a
}
