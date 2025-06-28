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
func SetupListHandlers(app *tview.Application, text *tview.InputField, list *tview.List, mainLayout tview.Primitive) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Handle key events for list widget
		if event.Key() == tcell.KeyTab {
			app.SetFocus(text)
			return nil
		}

		if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
			// Remove the currently selected item
			addRemoveItem(list)
			return nil
		}

		if event.Modifiers() == tcell.ModShift && event.Key() == tcell.KeyUp {
			// Implement logic to move selected item up
			moveSelectedItem(list, "up")

			return nil
		}

		if event.Modifiers() == tcell.ModShift && event.Key() == tcell.KeyDown {
			// Implement logic to move selected item down
			moveSelectedItem(list, "down")

			return nil
		}

		if event.Rune() == '?' {
			// Show help when user presses '?' button
			showHelpHandler(app, list, mainLayout)

			return nil
		}

		if event.Rune() == 'q' {
			app.Stop()

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

// Function for removing selected item from list
func addRemoveItem(list *tview.List) {
	list.RemoveItem(list.GetCurrentItem())

	items := GetTodoItems(list)

	if err := SaveTodoList(items); err != nil {
		fmt.Println("Error saving todo items:", err.Error())
	}
}

// Display generic modal dialog, can be used to display messages
func ShowModalDialog(app *tview.Application, parent tview.Primitive) {
	// Create the modal dialog
	modal := CreateModalDialog("Hello world!")

	// Setup modal handlers
	SetupModalHandlers(app, modal, parent)

	// Display it
	app.SetRoot(modal, true)
}

// Modal dialog handlers
func SetupModalHandlers(app *tview.Application, modal *tview.Modal, parent tview.Primitive) {
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		app.SetRoot(parent, true)
	})
}

// Function to move selected item up or down in the list
func moveSelectedItem(list *tview.List, direction string) {
	currentIndex := list.GetCurrentItem()
	itemCount := list.GetItemCount()

	switch direction {
	case "up":
		if currentIndex > 0 {
			// Get the current item's text
			currMainText, currSecondaryText := list.GetItemText(currentIndex)

			// Remove the current item
			list.RemoveItem(currentIndex)

			// Insert it at the previous position
			list.InsertItem(currentIndex-1, currMainText, currSecondaryText, '-', nil)

			// Set focus back to the moved item
			list.SetCurrentItem(currentIndex - 1)

			// Save the updated list
			items := GetTodoItems(list)

			if err := SaveTodoList(items); err != nil {
				fmt.Println("Error saving todo items:", err.Error())
			}
		}

	case "down":
		if currentIndex < itemCount-1 {
			// Get the current item's text
			mainText, secondaryText := list.GetItemText(currentIndex)

			// Remove the current item
			list.RemoveItem(currentIndex)

			// Insert it at the next position
			list.InsertItem(currentIndex+1, mainText, secondaryText, '-', nil)

			// Set focus back to the moved item
			list.SetCurrentItem(currentIndex + 1)

			// Save the updated list
			items := GetTodoItems(list)

			if err := SaveTodoList(items); err != nil {
				fmt.Println("Error saving todo items:", err.Error())
			}
		}
	}
}

// Function to handle '?' key press in list widget
func showHelpHandler(app *tview.Application, list *tview.List, mainLayout tview.Primitive) {
	// Create a help modal dialog
	modal := CreateModalDialog(`Help 
		
		Press Tab to toggle insert mode
		Backspace to delete items
		Shift + Up/Down to move items`)

	// Setup modal handlers to return to the main layout
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		// Return to the main layout and set focus to the list
		app.SetRoot(mainLayout, true).SetFocus(list)
	})

	// Display the modal
	app.SetRoot(modal, true)
}
