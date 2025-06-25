package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// SetupInputHandlers configures all input-related event handlers
func SetupInputHandlers(app *tview.Application, text *tview.InputField, list *tview.List) {
	// Enter key handler
	text.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			handleAddTodoItem(text, list)
		}
	})

	// Tab key handler for input field
	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(list)
			return nil
		}
		return event
	})
}

// SetupListHandlers configures all list-related event handlers
func SetupListHandlers(app *tview.Application, text *tview.InputField, list *tview.List) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(text)
			return nil
		}
		return event
	})
}

// handleAddTodoItem handles the business logic for adding a new todo item
func handleAddTodoItem(text *tview.InputField, list *tview.List) {
	newItem := text.GetText()
	if newItem != "" {
		list.AddItem(newItem, "", '-', nil)
		text.SetText("")
	}
}
