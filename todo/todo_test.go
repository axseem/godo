package todo

import (
	"reflect"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tl := &TaskList{}
	label := "task"
	tl.Add(label)

	want := &TaskList{{
		Label:     label,
		IsDone:    false,
		CreatedAt: (*tl)[0].CreatedAt,
		DoneAt:    time.Time{},
	}}

	if !reflect.DeepEqual(tl, want) {
		t.Errorf("error while adding task, got: %v, want: %v", tl, want)
	}
}

func TestRemove(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
	}

	tl.Remove(1, 4)

	if len(*tl) != len(labels)-2 {
		t.Errorf("wrong length, got: %v, want: %v", len(*tl), len(labels))
	}

	wantLabels := []string{"task1", "task3", "task4"}
	for i := 0; i < len(wantLabels); i++ {
		want := wantLabels[i]
		got := (*tl)[i].Label

		if got != want {
			t.Errorf("wrong labels, got: %v, want: %v", got, want)
		}
	}
}

func TestEdit(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
	}

	want := "edited"
	tl.Edit(1, want)
	got := (*tl)[1].Label

	if got != want {
		t.Errorf("wrong labels, got: %v, want: %v", got, want)
	}
}

func TestDone(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
	}

	err := tl.Done(1, 4)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	wantStates := []bool{false, true, false, false, true}
	for i := 0; i < len(wantStates); i++ {
		want := wantStates[i]
		got := (*tl)[i].IsDone

		if got != want {
			t.Errorf("wrong states, got: %v, want: %v", got, want)
		}
	}
}

func TestUndone(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
		tl.Done(i)
	}

	err := tl.Undone(1, 4)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	wantStates := []bool{true, false, true, true, false}
	for i := 0; i < len(wantStates); i++ {
		want := wantStates[i]
		got := (*tl)[i].IsDone

		if got != want {
			t.Errorf("wrong states, got: %v, want: %v", got, want)
		}
	}
}

func TestClear(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
	}

	tl.Done(1, 4)
	tl.Clear()

	wantLabels := []string{"task1", "task3", "task4"}
	for i := 0; i < len(wantLabels); i++ {
		want := wantLabels[i]
		got := (*tl)[i]

		if got.Label != want {
			t.Errorf("wrong labels, got: %v, want: %v", got.Label, want)
		}
		if got.IsDone != false {
			t.Errorf("wrong labels, got: %v, want: %v", got.IsDone, false)
		}
	}
}

func TestWipe(t *testing.T) {
	tl := &TaskList{}
	labels := []string{"task1", "task2", "task3", "task4", "task5"}
	for i := 0; i < len(labels); i++ {
		tl.Add(labels[i])
	}

	tl.Wipe()

	if *tl != nil {
		t.Errorf("wrong states, got: %v, want: %v", *tl, nil)
	}
}
