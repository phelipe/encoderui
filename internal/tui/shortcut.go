package tui

import (
	"github.com/gdamore/tcell/v2"
)

func (tui *TUI) ConfigureShortcuts() {
	tui.configureRuneShortcuts()
}

func (tui *TUI) configureRuneShortcuts() {

	tui.mainPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyEsc:
			tui.app.SetFocus(tui.mainPage)
			return nil
		}

		if event.Rune() != 0 && !tui.inputPanel.HasFocus() {

			switch event.Rune() {
			case 'E':
				tui.encode = true
				tui.outputPanel.SetText("Cliquei no encode", false)
				return nil
			case 'D':
				tui.encode = false
				tui.outputPanel.SetText("Cliquei no DEcode", false)
				return nil
			case 'A':
				tui.app.SetFocus(tui.encoderSelector)
				tui.outputPanel.SetText("Escolhe o encode a ser utilizado", false)
				return nil
			case 'I':
				tui.app.SetFocus(tui.inputPanel)
				return nil
			case 'O':
				tui.app.SetFocus(tui.outputPanel)
				return nil
			}
		}

		return event
	})
}
