package main

import (
	"github.com/phelipe/encoderui/internal/app"
	"github.com/phelipe/encoderui/internal/tui"
)

func main() {
	tui := tui.NewTUI(app.New())
	tui.CreatePage()
	tui.ConfigureShortcuts()
	if err := tui.Run(); err != nil {
		panic(err)
	}

}
