package cmd

import "fmt"

func Help() {
	var help string = `
godo is dead simple cli tool for managing your tasks.

Usage:

	godo <command> [arguments]

Commands:
	
	add [task]		add a task (u can add multiple tasks at once)
	clear			remove completed tasks
	do [index]		mark task as done
	edit [index][task]	edit task description
	ls 			list all tasks
	rm [index]		remove particular task
	undo [index]		mark task as not completed
	wipe			remove all tasks
	`

	fmt.Println(help)
}
