package cmd

import (
	"fmt"
	"os"

	"github.com/Tr1xN/godo/storage"
	"github.com/Tr1xN/godo/todo"
)

func Add(tl *todo.TaskList) {
	if len(os.Args) < 2 {
		fmt.Println("no task provided!")
		os.Exit(1)
	}
	for _, label := range os.Args[2:] {
		tl.Add(label)
		if err := storage.Save("db.json", tl); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
