package store

import (
	"reflect"
	"testing"

	"github.com/toversus/tbdist/model"
)

var getTasksTests = []struct {
	name   string
	ds     *Datastore
	status string
	expect []model.Task
	err    error
}{
	{
		name: "should return only pending tasks",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "go to school", Status: "DONE"},
				{ID: 2, Title: "withdraw my money", Status: "PENDING"},
				{ID: 3, Title: "play piano", Status: "DOING"},
				{ID: 4, Title: "go shopping", Status: "PENDING"},
			},
		},
		status: "PENDING",
		expect: []model.Task{
			{ID: 2, Title: "withdraw my money", Status: "PENDING"},
			{ID: 4, Title: "go shopping", Status: "PENDING"},
		},
	},
	{
		name: "should return only completed task",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "go to school", Status: "DONE"},
				{ID: 2, Title: "withdraw my money", Status: "PENDING"},
				{ID: 3, Title: "play piano", Status: "DOING"},
				{ID: 4, Title: "go shopping", Status: "PENDING"},
			},
		},
		status: "DONE",
		expect: []model.Task{
			{ID: 1, Title: "go to school", Status: "DONE"},
		},
	},
	{
		name: "should return only tasks in progress",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "go to school", Status: "DONE"},
				{ID: 2, Title: "withdraw my money", Status: "PENDING"},
				{ID: 3, Title: "play piano", Status: "DOING"},
				{ID: 4, Title: "go shopping", Status: "PENDING"},
			},
		},
		status: "DOING",
		expect: []model.Task{
			{ID: 3, Title: "play piano", Status: "DOING"},
		},
	},
}

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

func TestGetTasks(t *testing.T) {
	t.Log("getting tasks...")

	for _, testcase := range getTasksTests {
		t.Log(testcase.name)
		t.Log(testcase.status)
		if !reflect.DeepEqual(testcase.ds.getTasks(testcase.status), testcase.expect) {
			t.Errorf("=> Got %#v expected %#v", testcase.ds.tasks, testcase.expect)
		}
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
