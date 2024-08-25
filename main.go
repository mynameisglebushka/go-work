package main

import (
	"log"
)

func main() {
	var (
		application = newApplication()
		mainPage    = newMainPage(application)
		timerPage   = newTimerPage(application)
	)

	application.
		addMainPage(mainPage).
		addTimerPage(timerPage)

	if err := application.run(); err != nil {
		log.Fatal(err)
	}
}
