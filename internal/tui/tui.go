package tui

import (
	"github.com/phelipe/encoderui/internal/app"
	"github.com/rivo/tview"
)

type TUI struct {
	//use cases
	useCase *app.Application

	// app
	app *tview.Application

	// interation components
	inputPanel      *tview.TextArea
	outputPanel     *tview.TextArea
	encoderSelector *tview.DropDown

	//state
	encode bool

	//pages
	mainPage *tview.Flex
}

func NewTUI(app *app.Application) *TUI {
	return &TUI{
		useCase: app,

		app:             tview.NewApplication(),
		inputPanel:      tview.NewTextArea(),
		outputPanel:     tview.NewTextArea(),
		encoderSelector: tview.NewDropDown(),
		mainPage:        tview.NewFlex(),
	}
}
