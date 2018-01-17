package store

import (
	"reflect"
	"testing"
)

var saveTaskTests = []struct {
	name   string
	ds     *Datastore
	task   Task
	expect []Task
}{
	{
		name: "should save the new task in the datastore",
		ds:   &Datastore{},
		task: Task{Title: "withdraw my money", Status: "DOING"},
		expect: []Task{
			{1, "withdraw my money", "DOING"},
		},
	},
	{
		name: "should update the existing task in the datastore",
		ds: &Datastore{
			tasks: []Task{
				{1, "withdraw my money", "DOING"},
			},
		},
		task: Task{1, "withdraw my money", "DONE"},
		expect: []Task{
			{1, "withdraw my money", "DONE"},
		},
	},
}

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

func TestSaveTask(t *testing.T) {
	t.Log("saving task...")

	for _, testcase := range saveTaskTests {
		t.Log(testcase.name)
		t.Log(testcase.task.Status)
		testcase.ds.SaveTask(testcase.task)

		if !reflect.DeepEqual(testcase.ds.tasks, testcase.expect) {
			t.Errorf("=> Got %#v expected %#v", testcase.ds.tasks, testcase.expect)
		}
	}
}
