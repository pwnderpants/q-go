package cmd

import "github.com/rivo/tview"

func Execute() {
	// Create widgets
	list := CreateTodoList()
	text := CreateInputField()

	// Create application
	app := tview.NewApplication()

	// Setup handlers
	SetupInputHandlers(app, text, list)
	SetupListHandlers(app, text, list)

	// Create layout
	flex := CreateMainLayout(text, list)

	// Run application
	if err := app.SetRoot(flex, true).SetFocus(text).Run(); err != nil {
		panic(err)
	}
}
