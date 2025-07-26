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
	sidebar := CreateSubjectSidebar()

	// Create application
	app := tview.NewApplication()

	// Create layout
	flex := CreateMainLayout(text, list, sidebar)

	// Load saved app data on startup
	appData, err := LoadAppData()

	// If file doesn't exist just warn, but continue
	if err != nil {
		fmt.Println("Could not load app data: ", err)
		fmt.Println("Starting with default data.")

		appData = &AppData{
			Subjects:       []Subject{{Name: "General", Items: []TodoItem{}}},
			CurrentSubject: "General",
		}
	}

	// Load subjects into sidebar
	LoadSubjects(sidebar, appData.Subjects, appData.CurrentSubject)

	// Load current subject's todo items
	currentSubject := GetCurrentSubject(appData)

	if currentSubject != nil {
		LoadTodoItems(list, currentSubject.Items)
	}

	// Update todo list title to show current subject
	list.SetTitle("Todo list - " + appData.CurrentSubject + " (Press ? for help)")

	// Setup handlers
	SetupInputHandlers(app, text, list, sidebar, appData)
	SetupListHandlers(app, text, list, sidebar, flex, appData)
	SetupSidebarHandlers(app, sidebar, text, list, flex, appData)

	// Setup graceful shutdown
	setupSignalTrapper(app)

	// Run application
	if err := app.SetRoot(flex, true).SetFocus(sidebar).Run(); err != nil {
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
