package storage

import (
	"encoding/json"
	"os"

	"github.com/axsbee/godo/todo"
)

func Save(fileName string, tl *todo.TaskList) error {
	data, err := json.Marshal(tl)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func Load(fileName string, tl *todo.TaskList) error {
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			return err
		}
	}
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, tl)
}
