package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Execute() {
	// Create a listbox with sample items for main content
	category := tview.NewList().
		AddItem("Category", "", '1', nil).
		AddItem("Category", "", '2', nil).
		AddItem("Category", "", '3', nil).
		AddItem("Category", "", '4', nil).
		AddItem("Category", "", '5', nil).
		AddItem("Category", "", '6', nil).
		AddItem("Category", "", '7', nil).
		AddItem("Category", "", '8', nil)

	category.SetBorder(true)
	category.SetBackgroundColor(tcell.ColorDefault)
	category.SetMainTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	category.SetSecondaryTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))

	list := tview.NewList()

	// List of todo items
	list.SetTitle("Todo list (Press Tab to switch between input and list)")
	list.SetBorder(true)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetMainTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	list.SetSecondaryTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))

	// Input field
	text := tview.NewInputField().SetLabel("Add new todo item: ")

	text.SetBackgroundColor(tcell.ColorDefault)
	text.SetLabelColor(tcell.ColorDefault)
	text.SetLabelStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	text.SetFieldTextColor(tcell.ColorYellow)
	text.SetFieldBackgroundColor(tcell.ColorDefault)

	app := tview.NewApplication()

	// Set the function to run when Enter is pressed
	text.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			newItem := text.GetText()

			if newItem != "" {
				list.AddItem(newItem, "", '-', nil)
				text.SetText("")
			}
		}
	})

	// Handle Tab key to switch focus from input field to list
	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(list)
			return nil
		}
		return event
	})

	// Handle Tab key to switch focus from list to input field
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			app.SetFocus(text)
			return nil
		}
		return event
	})

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(text, 0, 1, true).
			AddItem(list, 0, 30, false), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(text).Run(); err != nil {
		panic(err)
	}
}
