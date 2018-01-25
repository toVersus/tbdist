package store

import (
	"errors"

	"github.com/toversus/tbdist/model"
)

// ErrTaskNotFound is returned when a Task ID is not found
var ErrTaskNotFound = errors.New("Task was not found")

// Datastore manages a list of tasks stored in memory
type Datastore struct {
	tasks  []model.Task
	lastID int // lastID is incremented for each new stored task
}

func (ds *Datastore) getTasks(status string) []model.Task {
	var tasks []model.Task
	for _, task := range ds.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetPendingTasks returns all the tasks putting on hold for now
func (ds *Datastore) GetPendingTasks() []model.Task {
	return ds.getTasks("PENDING")
}

// GetDoingTasks returns all the tasks in progress
func (ds *Datastore) GetDoingTasks() []model.Task {
	return ds.getTasks("DOING")
}

// GetDoneTasks returns all the completed tasks
func (ds *Datastore) GetDoneTasks() []model.Task {
	return ds.getTasks("DONE")
}

// SaveTask should save the task in the datastore if the task
// does not exist else update it. A Task Not Found error is returned
// when the task ID does not exist
func (ds *Datastore) SaveTask(task model.Task) error {
	if task.ID == 0 {
		ds.lastID++
		task.ID = ds.lastID
		ds.tasks = append(ds.tasks, task)
		return nil
	}

	for i, t := range ds.tasks {
		if t.ID == task.ID {
			ds.tasks[i] = task
			return nil
		}
	}

	return ErrTaskNotFound
}
