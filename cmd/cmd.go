package cmd

import (
	"fmt"

	"github.com/rivo/tview"
)

// To be passed to the main package for execution
func Execute() {
	// Create widgets
	list := CreateTodoList()
	text := CreateInputField()

	// Create application
	app := tview.NewApplication()

	// Create layout
	flex := CreateMainLayout(text, list)

	// Load saved todo items on startup
	items, err := LoadTodoList()

	// If file doesn't exist just warn, but continue
	if err != nil {
		fmt.Println("Could not load todo items: ", err)
		fmt.Println("Starting with an empty list.")
	} else {
		LoadTodoItems(list, items)
	}

	// Setup handlers
	SetupInputHandlers(app, text, list)
	SetupListHandlers(app, text, list)

	// Run application
	if err := app.SetRoot(flex, true).SetFocus(text).Run(); err != nil {
		panic(err)
	}
}
