package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/toversus/tbdist/store"
)

// Store defines the datastore services
type Store interface {
	GetPendingTasks() []store.Task
	SaveTask(task store.Task) error
}

var ds Store = &store.Datastore{}

type mockedStore struct {
	SaveTaskFunc func(task store.Task) error
}

func (ms *mockedStore) GetPendingTasks() []store.Task {
	return []store.Task{
		{1, "go to school", "PENDING"},
		{2, "withdraw my money", "PENDING"},
	}
}

func (ms *mockedStore) SaveTask(task store.Task) error {
	if ms.SaveTaskFunc != nil {
		return ms.SaveTaskFunc(task)
	}
	return nil
}

// GetPendingTasks returns pending tasks as a JSON response
func GetPendingTasks(w http.ResponseWriter, r *http.Request) {
	t := ds.GetPendingTasks()

	j, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// AddTask handles POST requests on /tasks.
// Return 201 if the task could be created
// Return 400 when JSON could not be decoded into a task
// datastore returned an error
func AddTask(w http.ResponseWriter, r *http.Request) {
	var t store.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ds.SaveTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateTask handles requests for updating an existing task,
// Return 200 if the task could be modified
// Return 400 when JSON could not be decoded into a task or
// datastore returned an error or task title is empty
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var t store.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ds.SaveTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateTask(t store.Task) error {
	if t.Title == "" {
		return errors.New("Title is missing")
	}
	return nil
}
