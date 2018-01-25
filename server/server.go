package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/toversus/tbdist/model"
	"github.com/toversus/tbdist/store"
)

// Store defines the datastore services
type Store interface {
	GetPendingTasks() []model.Task
	GetDoingTasks() []model.Task
	GetDoneTasks() []model.Task
	SaveTask(task model.Task) error
}

var ds Store = &store.Datastore{}

type mockedStore struct {
	SaveTaskFunc func(task model.Task) error
}

func (ms *mockedStore) GetPendingTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "PENDING"},
		{ID: 2, Title: "withdraw my money", Status: "PENDING"},
	}
}

func (ms *mockedStore) GetDoingTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "DOING"},
		{ID: 2, Title: "withdraw my money", Status: "DOING"},
	}
}

func (ms *mockedStore) GetDoneTasks() []model.Task {
	return []model.Task{
		{ID: 1, Title: "go to school", Status: "DONE"},
		{ID: 2, Title: "withdraw my money", Status: "DONE"},
	}
}

func (ms *mockedStore) SaveTask(task model.Task) error {
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

// GetDoingTasks returns tasks in progress as a JSON response
func GetDoingTasks(w http.ResponseWriter, r *http.Request) {
	t := ds.GetDoingTasks()
	j, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetDoneTasks returns tasks in progress as a JSON response
func GetDoneTasks(w http.ResponseWriter, r *http.Request) {
	t := ds.GetDoneTasks()
	j, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// AddTask handles POST requests on /tasks.
// Return 201 if the task could be created
// Return 400 when JSON could not be decoded into a task
// datastore returned an error
func AddTask(w http.ResponseWriter, r *http.Request) {
	var t model.Task

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
	var t model.Task

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

func validateTask(t model.Task) error {
	if t.Title == "" {
		return errors.New("Title is missing")
	}
	if t.Status != "DONE" && t.Status != "DOING" && t.Status != "PENDING" {
		return errors.New("Invalid status")
	}
	return nil
}
