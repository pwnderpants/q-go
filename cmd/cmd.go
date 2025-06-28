package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	SetupListHandlers(app, text, list, flex)

	// Setup graceful shutdown
	setupSignalTrapper(app)

	// Run application
	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		fmt.Printf("Application error: %v\n", err)
		os.Exit(1)
	}
}

// Setup graceful shutdown even with Ctrl+C or SIGTERM
func setupSignalTrapper(app *tview.Application) {
	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Run a goroutine to listen for signals
	go func() {
		<-sigChan

		app.Stop() // Gracefully stop the tview application
	}()
}
