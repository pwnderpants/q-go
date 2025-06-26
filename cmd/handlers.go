package cmd

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Handler for text input widget
func SetupInputHandlers(app *tview.Application, text *tview.InputField, list *tview.List) {

	// When user presses the enteer add text to list
	text.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			addTodoItem(text, list)
		}
	})

	// When user presses the tab key, switch focus to the list
	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(list)

			return nil
		}

		return event
	})
}

// Handler for list widget
func SetupListHandlers(app *tview.Application, text *tview.InputField, list *tview.List) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Handle key events for list widget
		if event.Key() == tcell.KeyTab {
			app.SetFocus(text)
			return nil
		}

		if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
			// Placeholder for delete functionality
			return nil
		}

		return event
	})
}

// Function for adding text to list
func addTodoItem(text *tview.InputField, list *tview.List) {
	newItem := text.GetText()

	if newItem != "" {
		list.AddItem(newItem, "", '-', nil)
		text.SetText("")

		items := GetTodoItems(list)

		if err := SaveTodoList(items); err != nil {
			fmt.Println("Error saving todo items:", err.Error())
		}
	}
}
