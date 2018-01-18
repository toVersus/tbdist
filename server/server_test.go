package server

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toversus/tbdist/store"
)

func TestGetPendingTasks(t *testing.T) {
	t.Log("getting pending tasks...")

	t.Log("should return pending tasks as JSON")

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/pending", nil)

	// The datastore is restored at the end of the test
	defer func() { ds = &store.Datastore{} }()

	ds = &mockedStore{}

	GetPendingTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("KO => Got %d expected %d", rec.Code, http.StatusOK)
	}

	expect := "[{\"id\":1,\"title\":\"go to school\",\"status\":\"PENDING\"},{\"id\":2,\"title\":\"withdraw my money\",\"status\":\"PENDING\"}]"
	if result := rec.Body.String(); result != expect {
		t.Errorf("KO => Got %s expected %s", result, expect)
	}
}

var addTaskTests = []struct {
	name     string
	saveFunc func(task store.Task) error
	body     []byte
	expect   int
}{
	{
		name:   "should add new task from JSON",
		body:   []byte(`{"Title":"buy bread for breakfast."}`),
		expect: http.StatusCreated,
	},
	{
		name:   "should return bad argument when JSON could not be handled",
		body:   []byte(""),
		expect: http.StatusBadRequest,
	},
	{
		name: "should response bad argument when datastore returns an error",
		saveFunc: func(task store.Task) error {
			return errors.New("datastore error")
		},
		body:   []byte(`["Title":"buy bread for breakfast."]`),
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
