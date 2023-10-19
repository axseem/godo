package cmd

import (
	"fmt"

	"github.com/axseem/godo/todo"
)

func checkmark(isDone bool) string {
	if isDone {
		return "☑"
	} else {
		return "☐"
	}
}

func taskListToString(tl *todo.TaskList) string {
	if len(*tl) == 0 {
		return "\tTask list is empty\n"
	}

	var output string
	for i, task := range *tl {
		output += fmt.Sprintf("\t%s %d. %s\n", checkmark(task.IsDone), i+1, task.Label)
	}
	return output
}

func List(tl *todo.TaskList) {
	fmt.Print(taskListToString(tl))
}
