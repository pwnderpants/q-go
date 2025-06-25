package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateTodoList creates and configures the todo list widget
func CreateTodoList() *tview.List {
	list := tview.NewList()
	list.SetTitle("Todo list (Press Tab to switch between input and list)")
	list.SetBorder(true)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetMainTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	list.SetSecondaryTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))

	return list
}

// CreateInputField creates and configures the input field widget
func CreateInputField() *tview.InputField {
	text := tview.NewInputField().SetLabel("Add new todo item: ")
	text.SetBackgroundColor(tcell.ColorDefault)
	text.SetLabelColor(tcell.ColorDefault)
	text.SetLabelStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	text.SetFieldTextColor(tcell.ColorYellow)
	text.SetFieldBackgroundColor(tcell.ColorDefault)

	return text
}
