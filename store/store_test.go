package store

import (
	"reflect"
	"testing"
)

func TestGetPendingTasks(t *testing.T) {
	t.Log("getting pending tasks...")

	ds := Datastore{
		tasks: []Task{
			{1, "go to school", "DONE"},
			{2, "withdraw my money", "PENDING"},
		},
	}

	expect := []Task{ds.tasks[1]}
	t.Log("should return the tasks witch need to be completed")
	if result := ds.GetPendingTasks(); !reflect.DeepEqual(result, expect) {
		t.Errorf("Got %#v expected %#v", result, expect)
	}
}

func TestSaveNewTask(t *testing.T) {
	t.Log("saving task...")

	ds := Datastore{}
	task := Task{Title: "withdraw my money"}
	expect := []Task{
		{1, "withdraw my money", "DOING"},
	}

	t.Log("should save the new task in the store")
	ds.SaveTask(task)
	if !reflect.DeepEqual(ds.tasks, expect) {
		t.Errorf("=> Got %#v expected %#v", ds.tasks, expect)
	}
}
