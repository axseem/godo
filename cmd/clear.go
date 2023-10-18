package cmd

import (
	"fmt"
	"os"

	"github.com/axsbee/godo/storage"
	"github.com/axsbee/godo/todo"
)

func Clear(tl *todo.TaskList) {
	tl.Clear()
	if err := storage.Save("db.json", tl); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
