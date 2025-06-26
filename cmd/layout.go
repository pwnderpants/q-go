package cmd

import "github.com/rivo/tview"

// CreateMainLayout creates the main application layout
func CreateMainLayout(text *tview.InputField, list *tview.List) *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(text, 3, 0, true).
			AddItem(list, 0, 1, false), 0, 2, false)
}
