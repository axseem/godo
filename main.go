package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/axseem/godo/cmd"
	"github.com/axseem/godo/storage"
	"github.com/axseem/godo/todo"
)

func main() {
	tl := &todo.TaskList{}
	flag.Parse()

	if err := storage.Load("db.json", tl); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "add":
		cmd.Add(tl)
		cmd.List(tl)
	case "rm", "remove", "del", "delete":
		cmd.Remove(tl)
		cmd.List(tl)
	case "edit":
		cmd.Edit(tl)
		cmd.List(tl)
	case "do", "done", "complete":
		cmd.Done(tl)
		cmd.List(tl)
	case "undone", "undo":
		cmd.Undone(tl)
		cmd.List(tl)
	case "ls", "list":
		cmd.List(tl)
	case "clear":
		cmd.Clear(tl)
		cmd.List(tl)
	case "wipe", "reset":
		cmd.Wipe(tl)
		cmd.List(tl)
	case "h", "help":
		cmd.Help()
	case "":
		cmd.List(tl)
	default:
		fmt.Println("Wrong command! Use 'help' command to see all avaiable commands")
	}
}
