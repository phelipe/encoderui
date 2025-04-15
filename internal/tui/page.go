package tui

import (
	"github.com/rivo/tview"
)

func (tui *TUI) createHeaderBarLayout() *tview.TextView {

	headerBar := tview.NewTextView().
		SetText("EncoderUI").
		SetTextAlign(tview.AlignCenter).SetScrollable(false).SetWrap(false)
	headerBar.SetBorder(true)

	return headerBar
}

func (tui *TUI) createConfigBarLayout() *tview.Flex {
	tabText := tview.NewTextView().SetText("[E]ncode | [D]ecode")

	keys := make([]string, 0)

	for k, _ := range tui.useCase.EncoderList {
		keys = append(keys, k)
	}

	tui.encoderSelector.
		SetLabel("algorithm  \n\n ").
		SetOptions(keys, nil).
		SetCurrentOption(0)

	configBar := tview.NewFlex().
		AddItem(tabText, 0, 1, false).
		AddItem(tui.encoderSelector, 0, 1, false)

	configBar.SetBorder(true)

	return configBar
}

func (tui *TUI) createContentLayout() *tview.Flex {
	tui.outputPanel.SetDisabled(true)
	content := tview.NewFlex().
		AddItem(tui.inputPanel, -1, 1, false).
		AddItem(tui.outputPanel, -1, 1, false)

	tui.inputPanel.SetBorder(true).SetTitle("Input value")
	tui.outputPanel.SetBorder(true).SetTitle("Output value")

	return content
}

func (tui *TUI) CreatePage() {

	tui.mainPage.SetDirection(tview.FlexRow).
		AddItem(tui.createHeaderBarLayout(), 3, 0, false).
		AddItem(tui.createConfigBarLayout(), 3, 0, false).
		AddItem(tui.createContentLayout(), 0, 1, false)

	tui.app.SetRoot(tui.mainPage, true).EnableMouse(true)
}

func (tui *TUI) Run() error {

	if err := tui.app.Run(); err != nil {
		return err
	}
	return nil
}
