package tui

import (
	"github.com/phelipe/encoderui/internal/app"
	"github.com/rivo/tview"
)

type TUI struct {
	//use cases
	app *app.Aplication

	// interation
	inputPanel  *tview.TextArea
	outputPanel *tview.TextView
	selector    *tview.DropDown

	//state
	encode bool

	//pages
}

func NewTUI() *TUI {
	return &TUI{
		app:         app.New(),
		inputPanel:  tview.NewTextArea(),
		outputPanel: tview.NewTextView(),
		selector:    tview.NewDropDown(),
	}
}
