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
				{ID: 1, Title: "go to school", Status: "DONE", Priority: 1},
				{ID: 2, Title: "withdraw my money", Status: "PENDING", Priority: 1},
				{ID: 3, Title: "play piano", Status: "DOING", Priority: 10},
				{ID: 4, Title: "go shopping", Status: "PENDING", Priority: 10},
			},
		},
		status: "PENDING",
		expect: []model.Task{
			{ID: 2, Title: "withdraw my money", Status: "PENDING", Priority: 1},
			{ID: 4, Title: "go shopping", Status: "PENDING", Priority: 10},
		},
	},
	{
		name: "should return only completed task",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "go to school", Status: "DONE", Priority: 1},
				{ID: 2, Title: "withdraw my money", Status: "PENDING", Priority: 3},
				{ID: 3, Title: "play piano", Status: "DOING", Priority: 5},
				{ID: 4, Title: "go shopping", Status: "PENDING", Priority: 7},
			},
		},
		status: "DONE",
		expect: []model.Task{
			{ID: 1, Title: "go to school", Status: "DONE", Priority: 1},
		},
	},
	{
		name: "should return only tasks in progress",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "go to school", Status: "DONE", Priority: 1},
				{ID: 2, Title: "withdraw my money", Status: "PENDING", Priority: 3},
				{ID: 3, Title: "play piano", Status: "DOING", Priority: 5},
				{ID: 4, Title: "go shopping", Status: "PENDING", Priority: 7},
			},
		},
		status: "DOING",
		expect: []model.Task{
			{ID: 3, Title: "play piano", Status: "DOING", Priority: 5},
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
		task: model.Task{Title: "withdraw my money", Status: "DOING", Priority: 1},
		expect: []model.Task{
			{ID: 1, Title: "withdraw my money", Status: "DOING", Priority: 1},
		},
	},
	{
		name: "should update the existing task in the datastore",
		ds: &Datastore{
			tasks: []model.Task{
				{ID: 1, Title: "withdraw my money", Status: "DOING", Priority: 9},
			},
		},
		task: model.Task{ID: 1, Title: "withdraw my money", Status: "DONE", Priority: 9},
		expect: []model.Task{
			{ID: 1, Title: "withdraw my money", Status: "DONE", Priority: 9},
		},
	},
	{
		name: "should return an error when task ID does not exist",
		ds:   &Datastore{},
		task: model.Task{ID: 1, Title: "withdraw my money", Status: "DONE", Priority: 5},
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
