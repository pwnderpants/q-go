package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Our data structure for todo items
type TodoItem struct {
	Text      string `yaml:"text"`
	Completed bool   `yaml:"completed"`
}

// Subject with its own todo list
type Subject struct {
	Name  string     `yaml:"name"`
	Items []TodoItem `yaml:"items"`
}

// The entire application data
type AppData struct {
	Subjects       []Subject `yaml:"subjects"`
	CurrentSubject string    `yaml:"current_subject"`
}

// Setup the storage path for save file
func getStoragePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Failed to get home directory: %w", err)
	}

	// Create .q-go dir if it doesn't exist
	qGoDir := filepath.Join(homeDir, ".q-go")

	if err := os.MkdirAll(qGoDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create .q-go directory: %w", err)
	}

	return filepath.Join(qGoDir, "data.yaml"), nil
}

// Saves the application data to YAML file
func SaveAppData(appData *AppData) error {
	storagePath, err := getStoragePath()

	if err != nil {
		return err
	}

	data, err := yaml.Marshal(appData)
	if err != nil {
		return fmt.Errorf("failed to marshal app data: %w", err)
	}

	if err := os.WriteFile(storagePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Loads the application data from YAML file
func LoadAppData() (*AppData, error) {
	storagePath, err := getStoragePath()

	if err != nil {
		return nil, err
	}

	// Check if file exists
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		// File doesn't exist, return default with General subject
		return &AppData{
			Subjects:       []Subject{{Name: "General", Items: []TodoItem{}}},
			CurrentSubject: "General",
		}, nil
	}

	data, err := os.ReadFile(storagePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var appData AppData

	if err := yaml.Unmarshal(data, &appData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal app data: %w", err)
	}

	// Ensure we have at least one subject
	if len(appData.Subjects) == 0 {
		appData.Subjects = []Subject{{Name: "General", Items: []TodoItem{}}}
		appData.CurrentSubject = "General"
	}

	// Ensure current subject exists
	found := false
	for _, subject := range appData.Subjects {
		if subject.Name == appData.CurrentSubject {
			found = true
			break
		}
	}
	if !found {
		appData.CurrentSubject = appData.Subjects[0].Name
	}

	return &appData, nil
}
