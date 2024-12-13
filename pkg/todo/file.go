package todo

import (
	"encoding/json"
	"os"
)

const FileName = "tasks.json"

func (t *TodoList) LoadTasks() error {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return nil // No file, no tasks to load
	}
	data, err := os.ReadFile(FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &t.Tasks)
}

func (t *TodoList) SaveTasks() error {
	data, err := json.Marshal(t.Tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(FileName, data, 0644)
}
