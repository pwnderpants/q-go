package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Setup and configure list widget
func CreateTodoList() *tview.List {
	list := tview.NewList()
	list.SetTitle("Todo list (Press Tab to switch between input and list)")
	list.SetBorder(true)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetMainTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	list.SetSecondaryTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))

	return list
}

// Setup and configure input widget
func CreateInputField() *tview.InputField {
	text := tview.NewInputField().SetLabel("Add new todo item: ")
	text.SetBackgroundColor(tcell.ColorDefault)
	text.SetLabelColor(tcell.ColorDefault)
	text.SetLabelStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	text.SetFieldTextColor(tcell.ColorYellow)
	text.SetFieldBackgroundColor(tcell.ColorDefault)

	return text
}

// Get all current items on the list
func GetTodoItems(list *tview.List) []TodoItem {
	var items []TodoItem

	for i := 0; i < list.GetItemCount(); i++ {
		mainText, _ := list.GetItemText(i)

		items = append(items, TodoItem{
			Text:      mainText,
			Completed: false,
		})
	}

	return items
}

// Load all items to the list
func LoadTodoItems(list *tview.List, items []TodoItem) {
	list.Clear()

	for _, item := range items {
		list.AddItem(item.Text, "", '-', nil)
	}
}

// Create and configure modal dialog widget
func CreateModalDialog(msg string) *tview.Modal {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetBackgroundColor((tcell.ColorDefault)).
		SetTextColor((tcell.ColorRed))

	return modal
}
