// Package todo provides functionality for managing a task list.
package todo

import (
	"fmt"
	"sort"
	"time"
)

type task struct {
	Label     string
	IsDone    bool
	CreatedAt time.Time
	DoneAt    time.Time
}

// TaskList is a slice of task structs, used to represent a list of tasks.
type TaskList []task

// Appends a new task with the given label to the end of the task list.
// The new task is initialized with IsDone set to false
// and CreatedAt set to the current time.
func (tl *TaskList) Add(label string) {
	t := task{
		Label:     label,
		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    time.Time{},
	}

	*tl = append(*tl, t)
}

// Removes tasks from the list by index.
// The indexes are sorted in reverse order before removing
// so that removing doesn't impact indexes that come later.
// Returns an error if any index is out of range.
func (tl *TaskList) Remove(il ...int) error {
	sort.Sort(sort.Reverse(sort.IntSlice(il)))
	for _, i := range il {
		if i < 0 || i >= len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		*tl = append((*tl)[:i], (*tl)[i+1:]...)
	}

	return nil
}

// Updates the label of the task at index i in the TaskList.
// Returns an error if i is out of range.
func (tl *TaskList) Edit(i int, label string) error {
	if i < 0 || i >= len(*tl) {
		return fmt.Errorf("item %d dose not exist", i)
	}
	(*tl)[i].Label = label

	return nil
}

// Marks the tasks at the given indexes as done
// by setting IsDone to true and DoneAt to the current time.
// Returns an error if any index is out of range.
func (tl *TaskList) Done(il ...int) error {
	for _, i := range il {
		if i < 0 || i >= len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		(*tl)[i].IsDone = true
		(*tl)[i].DoneAt = time.Now()
	}

	return nil
}

// Marks the tasks at the given indexes as undone
// by setting IsDone to false and DoneAt to the zero time.
// Returns an error if any index is out of range.
func (tl *TaskList) Undone(il ...int) error {
	for _, i := range il {
		if i < 0 || i >= len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		(*tl)[i].IsDone = false
		(*tl)[i].DoneAt = time.Time{}
	}

	return nil
}

// Removes all completed tasks from the TaskList.
func (tl *TaskList) Clear() {
	var newTaskList TaskList
	for _, task := range *tl {
		if !task.IsDone {
			newTaskList = append(newTaskList, task)
		}
	}
	*tl = newTaskList
}

// Wipe removes all tasks from the TaskList by setting it to nil.
func (tl *TaskList) Wipe() {
	*tl = nil
}
