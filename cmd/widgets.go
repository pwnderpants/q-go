package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Setup and configure list widget
func CreateTodoList() *tview.List {
	list := tview.NewList()
	list.SetTitle("Todo list - (Press ? for help)")
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
	text.SetBorder(true)

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

// Create and configure sidebar for subjects
func CreateSubjectSidebar() *tview.List {
	sidebar := tview.NewList()
	sidebar.SetTitle("Subjects")
	sidebar.SetBorder(true)
	sidebar.SetBackgroundColor(tcell.ColorDefault)
	sidebar.SetMainTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	sidebar.SetSecondaryTextStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
	sidebar.ShowSecondaryText(false)
	sidebar.SetWrapAround(false)

	return sidebar
}

// Load subjects into sidebar
func LoadSubjects(sidebar *tview.List, subjects []Subject, currentSubject string) {
	sidebar.Clear()

	for i, subject := range subjects {
		prefix := "  "
		if subject.Name == currentSubject {
			prefix = "→ "
		}
		sidebar.AddItem(prefix+subject.Name, "", rune('1'+i), nil)
	}
}

// Get current subject from app data
func GetCurrentSubject(appData *AppData) *Subject {
	for i := range appData.Subjects {
		if appData.Subjects[i].Name == appData.CurrentSubject {
			return &appData.Subjects[i]
		}
	}
	return nil
}

// Add new subject to app data
func AddSubject(appData *AppData, name string) {
	appData.Subjects = append(appData.Subjects, Subject{Name: name, Items: []TodoItem{}})
}

// Delete subject from app data
func DeleteSubject(appData *AppData, name string) bool {
	if len(appData.Subjects) <= 1 {
		return false // Don't delete if it's the last subject
	}

	for i, subject := range appData.Subjects {
		if subject.Name == name {
			appData.Subjects = append(appData.Subjects[:i], appData.Subjects[i+1:]...)
			if appData.CurrentSubject == name {
				appData.CurrentSubject = appData.Subjects[0].Name
			}
			return true
		}
	}
	return false
}

// Rename subject in app data
func RenameSubject(appData *AppData, oldName, newName string) bool {
	if newName == "" || newName == oldName {
		return false
	}

	// Check if new name already exists
	for _, subject := range appData.Subjects {
		if subject.Name == newName {
			return false
		}
	}

	for i := range appData.Subjects {
		if appData.Subjects[i].Name == oldName {
			appData.Subjects[i].Name = newName
			if appData.CurrentSubject == oldName {
				appData.CurrentSubject = newName
			}
			return true
		}
	}
	return false
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

// Create input modal for new subject
func CreateInputModal(title, label string) *tview.Form {
	form := tview.NewForm()
	form.SetTitle(title)
	form.SetBorder(true)
	form.SetBackgroundColor(tcell.ColorDefault)
	form.AddInputField(label, "", 20, nil, nil)
	form.AddButton("OK", nil)
	form.AddButton("Cancel", nil)

	return form
}
