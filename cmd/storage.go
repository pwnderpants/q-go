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

// The entire todo list
type TodoList struct {
	Items []TodoItem `yaml:"todos"`
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

	return filepath.Join(qGoDir, "list.yaml"), nil
}

// Saves the todo list to YAML file
func SaveTodoList(items []TodoItem) error {
	storagePath, err := getStoragePath()

	if err != nil {
		return err
	}

	todoList := TodoList{Items: items}

	data, err := yaml.Marshal(todoList)
	if err != nil {
		return fmt.Errorf("failed to marshal todo list: %w", err)
	}

	if err := os.WriteFile(storagePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Loads the todo list from YAML file
func LoadTodoList() ([]TodoItem, error) {
	storagePath, err := getStoragePath()

	if err != nil {
		return nil, err
	}

	// Check if file exists
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		// File doesn't exist, return empty list
		return []TodoItem{}, nil
	}

	data, err := os.ReadFile(storagePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var todoList TodoList

	if err := yaml.Unmarshal(data, &todoList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal todo list: %w", err)
	}

	return todoList.Items, nil
}
