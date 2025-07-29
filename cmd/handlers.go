package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Handler for text input widget
func SetupInputHandlers(app *tview.Application, text *tview.InputField, list *tview.List, sidebar *tview.List, appData *AppData) {

	// When user presses the enteer add text to list
	text.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			addTodoItem(text, list, appData)
		}
	})

	// When user presses the tab key, switch focus to the list
	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(list)

			return nil
		}

		if event.Key() == tcell.KeyEsc {
			app.SetFocus(sidebar)

			return nil
		}

		return event
	})
}

// Handler for list widget
func SetupListHandlers(app *tview.Application, text *tview.InputField, list *tview.List, sidebar *tview.List, mainLayout tview.Primitive, appData *AppData) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Handle key events for list widget
		if event.Key() == tcell.KeyTab {
			app.SetFocus(sidebar)

			return nil
		}

		if event.Key() == tcell.KeyEsc {
			app.SetFocus(sidebar)

			return nil
		}

		if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
			// Remove the currently selected item
			removeItem(list, appData)

			return nil
		}

		if event.Modifiers() == tcell.ModShift && event.Key() == tcell.KeyUp {
			// Implement logic to move selected item up
			moveSelectedItem(list, "up", appData)

			return nil
		}

		if event.Modifiers() == tcell.ModShift && event.Key() == tcell.KeyDown {
			// Implement logic to move selected item down
			moveSelectedItem(list, "down", appData)

			return nil
		}

		if event.Rune() == 'e' {
			// Edit currently selected item
			editSelectedItem(app, list, mainLayout, appData)

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

// Handler for sidebar widget
func SetupSidebarHandlers(app *tview.Application, sidebar *tview.List, text *tview.InputField, list *tview.List, mainLayout tview.Primitive, appData *AppData) {
	sidebar.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(text)

			return nil
		}

		if event.Key() == tcell.KeyEnter {
			// Switch to selected subject
			currentIndex := sidebar.GetCurrentItem()

			if currentIndex >= 0 && currentIndex < len(appData.Subjects) {
				appData.CurrentSubject = appData.Subjects[currentIndex].Name

				LoadSubjects(sidebar, appData.Subjects, appData.CurrentSubject)
				refreshTodoList(list, appData)
				list.SetTitle("Todo list - " + appData.CurrentSubject + " (Press ? for help)")
				SaveAppData(appData)
			}

			return nil
		}

		if event.Rune() == 'n' {
			// Create new subject
			showNewSubjectModal(app, sidebar, list, mainLayout, appData)

			return nil
		}

		if event.Rune() == 'd' {
			// Delete current subject
			if len(appData.Subjects) > 1 {
				currentIndex := sidebar.GetCurrentItem()

				if currentIndex >= 0 && currentIndex < len(appData.Subjects) {
					subjectName := appData.Subjects[currentIndex].Name

					if DeleteSubject(appData, subjectName) {
						LoadSubjects(sidebar, appData.Subjects, appData.CurrentSubject)
						refreshTodoList(list, appData)
						list.SetTitle("Todo list - " + appData.CurrentSubject + " (Press ? for help)")
						SaveAppData(appData)
					}
				}
			}

			return nil
		}

		if event.Rune() == 'r' {
			// Rename current subject
			currentIndex := sidebar.GetCurrentItem()

			if currentIndex >= 0 && currentIndex < len(appData.Subjects) {
				subjectName := appData.Subjects[currentIndex].Name

				showRenameSubjectModal(app, sidebar, list, mainLayout, appData, subjectName)
			}

			return nil
		}

		if event.Rune() == '?' {
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
func addTodoItem(text *tview.InputField, list *tview.List, appData *AppData) {
	newItem := text.GetText()

	if newItem != "" {
		list.AddItem(newItem, "", '-', nil)
		text.SetText("")

		currentSubject := GetCurrentSubject(appData)

		if currentSubject != nil {
			currentSubject.Items = GetTodoItems(list)

			SaveAppData(appData)
		}
	}
}

// Function for removing selected item from list
func removeItem(list *tview.List, appData *AppData) {
	list.RemoveItem(list.GetCurrentItem())

	currentSubject := GetCurrentSubject(appData)

	if currentSubject != nil {
		currentSubject.Items = GetTodoItems(list)
		SaveAppData(appData)
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
func moveSelectedItem(list *tview.List, direction string, appData *AppData) {
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
			currentSubject := GetCurrentSubject(appData)
			if currentSubject != nil {
				currentSubject.Items = GetTodoItems(list)
				SaveAppData(appData)
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
			currentSubject := GetCurrentSubject(appData)

			if currentSubject != nil {
				currentSubject.Items = GetTodoItems(list)
				SaveAppData(appData)
			}
		}
	}
}

// Refresh todo list with current subject items
func refreshTodoList(list *tview.List, appData *AppData) {
	currentSubject := GetCurrentSubject(appData)

	if currentSubject != nil {
		LoadTodoItems(list, currentSubject.Items)
	}
}

// Show modal for creating new subject
func showNewSubjectModal(app *tview.Application, sidebar *tview.List, list *tview.List, mainLayout tview.Primitive, appData *AppData) {
	form := CreateInputModal("New Subject", "Subject name:")

	// Ensure Tab navigation works in the form
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Let the form handle Tab navigation naturally
		return event
	})

	form.GetButton(0).SetSelectedFunc(func() {
		// OK button
		subjectName := form.GetFormItem(0).(*tview.InputField).GetText()

		if subjectName != "" {
			AddSubject(appData, subjectName)
			appData.CurrentSubject = subjectName
			LoadSubjects(sidebar, appData.Subjects, appData.CurrentSubject)
			refreshTodoList(list, appData)
			list.SetTitle("Todo list - " + appData.CurrentSubject + " (Press ? for help)")
			SaveAppData(appData)
		}
		app.SetRoot(mainLayout, true).SetFocus(sidebar)
	})

	form.GetButton(1).SetSelectedFunc(func() {
		// Cancel button
		app.SetRoot(mainLayout, true).SetFocus(sidebar)
	})

	app.SetRoot(form, true)
}

// Show modal for editing selected item
func editSelectedItem(app *tview.Application, list *tview.List, mainLayout tview.Primitive, appData *AppData) {
	currentIndex := list.GetCurrentItem()
	if currentIndex < 0 || currentIndex >= list.GetItemCount() {
		return
	}

	currentText, _ := list.GetItemText(currentIndex)
	form := CreateInputModal("Edit Item", "Item text:")

	// Pre-fill with current text
	inputField := form.GetFormItem(0).(*tview.InputField)
	inputField.SetText(currentText)

	// Handle ESC key and allow Tab navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			// ESC acts like Cancel button
			app.SetRoot(mainLayout, true).SetFocus(list)

			return nil
		}
		// Let the form handle Tab navigation naturally
		return event
	})

	form.GetButton(0).SetSelectedFunc(func() {
		// OK button
		newText := inputField.GetText()

		if newText != "" {
			// Update the item text
			list.SetItemText(currentIndex, newText, "")

			// Save changes
			currentSubject := GetCurrentSubject(appData)
			if currentSubject != nil {
				currentSubject.Items = GetTodoItems(list)

				SaveAppData(appData)
			}
		}
		app.SetRoot(mainLayout, true).SetFocus(list)
	})

	form.GetButton(1).SetSelectedFunc(func() {
		// Cancel button
		app.SetRoot(mainLayout, true).SetFocus(list)
	})

	app.SetRoot(form, true)
}

// Show modal for renaming a subject
func showRenameSubjectModal(app *tview.Application, sidebar *tview.List, list *tview.List, mainLayout tview.Primitive, appData *AppData, oldName string) {
	form := CreateInputModal("Rename Subject", "New name:")

	// Pre-fill with current name
	inputField := form.GetFormItem(0).(*tview.InputField)

	inputField.SetText(oldName)

	// Handle ESC key and allow Tab navigation
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			// ESC acts like Cancel button
			app.SetRoot(mainLayout, true).SetFocus(sidebar)

			return nil
		}
		// Let the form handle Tab navigation naturally
		return event
	})

	form.GetButton(0).SetSelectedFunc(func() {
		// OK button
		newName := inputField.GetText()

		if newName != "" && RenameSubject(appData, oldName, newName) {
			LoadSubjects(sidebar, appData.Subjects, appData.CurrentSubject)
			refreshTodoList(list, appData)
			list.SetTitle("Todo list - " + appData.CurrentSubject + " (Press ? for help)")
			SaveAppData(appData)
		}
		app.SetRoot(mainLayout, true).SetFocus(sidebar)
	})

	form.GetButton(1).SetSelectedFunc(func() {
		// Cancel button
		app.SetRoot(mainLayout, true).SetFocus(sidebar)
	})

	app.SetRoot(form, true)
}

// Function to handle '?' key press in list widget
func showHelpHandler(app *tview.Application, list *tview.List, mainLayout tview.Primitive) {
	// Create a help modal dialog
	modal := CreateModalDialog(`Help - Navigation & Controls
		
		Navigation:
		Tab: Cycle focus through Input → Todo List → Subjects
		Esc: Focus Subjects panel from any pane
		
		Input Panel:
		Enter: Add todo item to current subject
		
		Todo List:
		'e': Edit selected item
		Backspace: Delete selected item
		Shift + Up/Down: Move items up/down
		
		Subjects Panel:
		Enter: Select/switch to subject
		'n': Create new subject
		'r': Rename current subject
		'd': Delete current subject
		
		General:
		'?': Show this help dialog
		'q': Quit application`)

	// Setup modal handlers to return to the main layout
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		// Return to the main layout and set focus to the list
		app.SetRoot(mainLayout, true).SetFocus(list)
	})

	// Display the modal
	app.SetRoot(modal, true)
}
