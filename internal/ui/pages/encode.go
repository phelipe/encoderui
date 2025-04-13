package pages

import (
	"github.com/phelipe/encoderui/internal/app"
	"github.com/rivo/tview"
)

type TUI struct {
	app *app.Aplication
}

func TUIRun() {
	app := tview.NewApplication()

	// header
	asciiArt := `
▄▖       ▌    ▖▖▄▖
▙▖▛▌▛▘▛▌▛▌█▌▛▘▌▌▐ 
▙▖▌▌▙▖▙▌▙▌▙▖▌ ▙▌▟▖
`

	text := tview.NewTextView().
		SetText(asciiArt).
		SetTextAlign(tview.AlignLeft).SetScrollable(false).SetWrap(false)
	// SetDynamicColors(false)

	text.SetBorder(false)

	selector := tview.NewDropDown().
		SetLabel("algorithm  \n\n ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil).
		SetCurrentOption(0)

	header := tview.NewFlex().
		AddItem(text, 0, 3, false).
		AddItem(selector, 0, 1, false)

	header.SetBorder(true)

	//main content
	inputBox := tview.NewTextArea().SetPlaceholder("texto aqui dfkjalsdfj aldkfjk")
	inputBox.SetBorder(true).SetTitle("Input value")
	outputBox := tview.NewTextView()
	outputBox.SetBorder(true).SetTitle("Output value")

	content := tview.NewFlex().
		AddItem(inputBox, 0, 1, false).
		AddItem(outputBox, 0, 1, false)

	//full page
	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 6, 0, false).
		AddItem(content, 0, 1, false)

	if err := app.SetRoot(page, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
