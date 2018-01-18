package server

import (
	"encoding/json"
	"net/http"

	"github.com/toversus/tbdist/store"
)

// Store defines the datastore services
type Store interface {
	GetPendingTasks() []store.Task
}

var ds Store = &store.Datastore{}

type mockedStore struct{}

func (ms *mockedStore) GetPendingTasks() []store.Task {
	return []store.Task{
		{1, "go to school", "PENDING"},
		{2, "withdraw my money", "PENDING"},
	}
}

// GetPendingTasks returns pending tasks as a JSON response
func GetPendingTasks(w http.ResponseWriter, r *http.Request) {
	t := ds.GetPendingTasks()

	j, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
