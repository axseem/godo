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

type TaskList []task

func (tl *TaskList) Add(label string) {
	t := task{
		Label:     label,
		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    time.Time{},
	}

	*tl = append(*tl, t)
}

func (tl *TaskList) Remove(il []int) error {
	sort.Sort(sort.Reverse(sort.IntSlice(il)))
	for _, i := range il {
		if i <= 0 || i > len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		*tl = append((*tl)[:i-1], (*tl)[i:]...)
	}

	return nil
}

func (tl *TaskList) Edit(i int, label string) error {
	if i <= 0 || i > len(*tl) {
		return fmt.Errorf("item %d dose not exist", i)
	}
	(*tl)[i-1].Label = label

	return nil
}

func (tl *TaskList) Done(il []int) error {
	for _, i := range il {
		if i <= 0 || i > len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		(*tl)[i-1].IsDone = true
		(*tl)[i-1].DoneAt = time.Now()
	}

	return nil
}

func (tl *TaskList) Undone(il []int) error {
	for _, i := range il {
		if i <= 0 || i > len(*tl) {
			return fmt.Errorf("item %d dose not exist", i)
		}
		(*tl)[i-1].IsDone = false
		(*tl)[i-1].DoneAt = time.Time{}
	}

	return nil
}

func (tl *TaskList) Clear() {
	var newTaskList TaskList
	for _, task := range *tl {
		if !task.IsDone {
			newTaskList = append(newTaskList, task)
		}
	}
	*tl = newTaskList
}

func (tl *TaskList) Wipe() {
	*tl = nil
}
