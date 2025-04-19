package cmd

import (
	"encoding/json"
	"os"
)

var todoFile = "todos.json"

func readTodos() []string {
	var todos []string
	file, err := os.ReadFile(todoFile)
	if err == nil {
		json.Unmarshal(file, &todos)
	}
	return todos
}

func saveTodos(todos []string) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile(todoFile, data, 0644)
}
