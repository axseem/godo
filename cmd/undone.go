package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/axseem/godo/storage"
	"github.com/axseem/godo/todo"
)

func Undone(tl *todo.TaskList) {
	if len(os.Args) < 2 {
		fmt.Println("no task index provided!")
		os.Exit(1)
	}
	var il []int
	for _, i := range os.Args[2:] {
		i, err := strconv.Atoi(i)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		il = append(il, i-1)
	}
	if err := tl.Undone(il); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := storage.Save("db.json", tl); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
