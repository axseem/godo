package cmd

import (
	"fmt"
	"os"

	"github.com/Tr1xN/godo/storage"
	"github.com/Tr1xN/godo/todo"
)

func Wipe(tl *todo.TaskList) {
	tl.Wipe()
	if err := storage.Save("db.json", tl); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
