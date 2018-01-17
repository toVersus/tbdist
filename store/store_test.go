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
