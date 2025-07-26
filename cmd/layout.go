package cmd

import "github.com/rivo/tview"

// CreateMainLayout creates the main application layout with sidebar
func CreateMainLayout(text *tview.InputField, list *tview.List, sidebar *tview.List) *tview.Flex {
	// Create the main content area (input + todo list)
	mainContent := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(text, 3, 0, false).
		AddItem(list, 0, 1, false)

	// Create the main layout with sidebar and content
	return tview.NewFlex().
		AddItem(sidebar, 20, 0, true).
		AddItem(mainContent, 0, 1, false)
}
