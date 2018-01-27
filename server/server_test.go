package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toversus/tbdist/model"
	"github.com/toversus/tbdist/store"
)

type mockedStore struct {
	GetTaskFunc  func() []model.Task
	SaveTaskFunc func(task model.Task) error
}

func (ms *mockedStore) GetPendingTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "PENDING", Priority: 1},
		{ID: 2, Title: "withdraw my money", Status: "PENDING", Priority: 10},
	}
}

func (ms *mockedStore) GetDoingTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "DOING", Priority: 1},
		{ID: 2, Title: "withdraw my money", Status: "DOING", Priority: 10},
	}
}

func (ms *mockedStore) GetDoneTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "DONE", Priority: 1},
		{ID: 2, Title: "withdraw my money", Status: "DONE", Priority: 10},
	}
}

func (ms *mockedStore) SaveTask(task model.Task) error {
	if ms.SaveTaskFunc != nil {
		return ms.SaveTaskFunc(task)
	}
	return nil
}

var getTaskTests = []struct {
	name    string
	getFunc func() []model.Task
	getTask func(w http.ResponseWriter, r *http.Request)
	expect  string
}{
	{
		name:    "should return pending tasks as JSON",
		getTask: GetPendingTasks,
		expect:  "[{\"id\":1,\"title\":\"go to school\",\"status\":\"PENDING\",\"priority\":1},{\"id\":2,\"title\":\"withdraw my money\",\"status\":\"PENDING\",\"priority\":10}]",
	},
	{
		name:    "should return tasks in progress as JSON",
		getTask: GetDoingTasks,
		expect:  "[{\"id\":1,\"title\":\"go to school\",\"status\":\"DOING\",\"priority\":1},{\"id\":2,\"title\":\"withdraw my money\",\"status\":\"DOING\",\"priority\":10}]",
	},
	{
		name:    "should return completed tasks as JSON",
		getTask: GetDoneTasks,
		expect:  "[{\"id\":1,\"title\":\"go to school\",\"status\":\"DONE\",\"priority\":1},{\"id\":2,\"title\":\"withdraw my money\",\"status\":\"DONE\",\"priority\":10}]",
	},
}

func TestGetPendingTasks(t *testing.T) {
	t.Log("getting pending tasks...")

	for _, testcase := range getTaskTests {
		t.Log(testcase.name)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/tasks/pending", nil)

		// The datastore is restored at the end of the test
		defer func() { ds = &store.Datastore{} }()

		ds = &mockedStore{
			GetTaskFunc: testcase.getFunc,
		}

		testcase.getTask(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("KO => Got %d expected %d", rec.Code, http.StatusOK)
		}
		if result := rec.Body.String(); result != testcase.expect {
			t.Errorf("KO => Got %s expected %s", result, testcase.expect)
		}
	}
}

var addTaskTests = []struct {
	name     string
	saveFunc func(task model.Task) error
	body     []byte
	expect   int
}{
	{
		name:   "should add new task from JSON",
		body:   []byte(`{"Title":"buy bread for breakfast.","Status":"DOING", "Priority":1}`),
		expect: http.StatusCreated,
	},
	{
		name:   "should return bad argument when JSON could not be handled",
		body:   []byte(""),
		expect: http.StatusBadRequest,
	},
	{
		name: "should response bad argument when datastore returns an error",
		saveFunc: func(task model.Task) error {
			return errors.New("datastore error")
		},
		body:   []byte(`["Title":"buy bread for breakfast."]`),
		expect: http.StatusBadRequest,
	},
	{
		name:   "should response bad argument when task title is empty",
		body:   []byte(`["Title":""]`),
		expect: http.StatusBadRequest,
	},
	{
		name:   "should response bad argument when task status is invalid",
		body:   []byte(`["Title":"buy bread for breakfast.","Status":"HOGE","Priority":1]`),
		expect: http.StatusBadRequest,
	},
	{
		name:   "should response bad argument when task priority is invalid",
		body:   []byte(`["Title":"buy bread for breakfast.","Status":"DOING","Priority":100]`),
		expect: http.StatusBadRequest,
	},
}

func TestAddTask(t *testing.T) {
	t.Log("adding task...")

	for _, testcase := range addTaskTests {
		t.Log(testcase.name)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(testcase.body))

		defer func() { ds = &store.Datastore{} }()

		ds = &mockedStore{
			SaveTaskFunc: testcase.saveFunc,
		}

		AddTask(rec, req)

		if rec.Code != testcase.expect {
			t.Errorf("KO => Got %d expected %d", rec.Code, testcase.expect)
		}
	}
}

var updateTaskTests = []struct {
	name     string
	saveFunc func(task model.Task) error
	body     []byte
	expect   int
}{
	{
		name:   "should response with a status 200 OK the task was updated",
		body:   []byte(`{"ID":1, "Title":"buy bread for breakfast.", "Status":"DONE", "Priority":1}`),
		expect: http.StatusOK,
	},
	{
		name:   "should response with a statu 400 Bad Request when JSON body could not be handled",
		body:   []byte(""),
		expect: http.StatusBadRequest,
	},
	{
		name: "should response with a statu 400 Bad Request when the datastore returned an error",
		saveFunc: func(task model.Task) error {
			return errors.New("datastore error")
		},
		body:   []byte(`{"ID":1, "Title":"buy bread for breakfast.", "Status": "DONE", "Priority":1}`),
		expect: http.StatusBadRequest,
	},
	{
		name:   "should response with a status 400 Bad Request when task title is empty",
		body:   []byte(`{"Title":""}`),
		expect: http.StatusBadRequest,
	},
}

func TestUpdateTask(t *testing.T) {
	t.Log("updating task...")

	for _, testcase := range updateTaskTests {
		t.Logf(testcase.name)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tasks/1", bytes.NewBuffer(testcase.body))

		defer func() { ds = &store.Datastore{} }()

		ds = &mockedStore{
			SaveTaskFunc: testcase.saveFunc,
		}

		UpdateTask(rec, req)

		if rec.Code != testcase.expect {
			t.Errorf("KO => Got %d expected %d", rec.Code, testcase.expect)
		}
	}
}
