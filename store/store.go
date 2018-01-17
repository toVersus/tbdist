package store

import "errors"

// ErrTaskNotFound is returned when a Task ID is not found
var ErrTaskNotFound = errors.New("Task was not found")

// Task is thing to be done or completed
type Task struct {
	ID     int
	Title  string
	Status string // DOING, PENDING, DONE
}

// Datastore manages a list of tasks stored in memory
type Datastore struct {
	tasks  []Task
	lastID int // lastID is incremented for each new stored task
}

// GetPendingTasks returns all the tasks which need to be done
func (ds *Datastore) GetPendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range ds.tasks {
		if task.Status == "PENDING" {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

// SaveTask should save the task in the datastore if the task
// does not exist else update it. A Task Not Found error is returned
// when the task ID does not exist
func (ds *Datastore) SaveTask(task Task) error {
	if task.ID == 0 {
		ds.lastID++
		task.ID = ds.lastID
		task.Status = "DOING"
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
