package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/axsbee/godo/storage"
	"github.com/axsbee/godo/todo"
)

func Edit(tl *todo.TaskList) {
	if len(os.Args) < 2 {
		fmt.Println("no task index provided!")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		fmt.Println("no task label provided!")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tl.Edit(i, strings.Join(os.Args[3:], " "))
	if err := storage.Save("db.json", tl); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
