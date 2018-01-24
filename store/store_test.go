package store

import (
	"reflect"
	"testing"

	"github.com/toversus/tbdist/model"
)

var saveTaskTests = []struct {
	name   string
	ds     *Datastore
	task   model.Task
	expect []model.Task
	err    error
}{
	{
		name: "should save the new task in the datastore",
		ds:   &Datastore{},
		task: model.Task{Title: "withdraw my money", Status: "DOING"},
		expect: []model.Task{
			{ID: 1, Title: "withdraw my money", Status: "DOING"},
		},
	},
	{
		name: "should update the existing task in the datastore",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "withdraw my money", Status: "DOING"},
			},
		},
		task: model.Task{ID: 1, Title: "withdraw my money", Status: "DONE"},
		expect: []model.Task{
			{ID: 1, Title: "withdraw my money", Status: "DONE"},
		},
	},
	{
		name: "should return an error when task ID does not exist",
		ds:   &Datastore{},
		task: model.Task{ID: 1, Title: "withdraw my money", Status: "DONE"},
		err:  ErrTaskNotFound,
	},
}

func TestGetPendingTasks(t *testing.T) {
	t.Log("getting pending tasks...")

	ds := Datastore{
		tasks: []model.Task{
			{ID: 1, Title: "go to school", Status: "DONE"},
			{ID: 2, Title: "withdraw my money", Status: "PENDING"},
		},
	}

	expect := []model.Task{ds.tasks[1]}
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
		t.Log(testcase.task.Status)
		if !reflect.DeepEqual(testcase.ds.tasks, testcase.expect) {
			t.Errorf("=> Got %#v expected %#v", testcase.ds.tasks, testcase.expect)
		}
	}
}
