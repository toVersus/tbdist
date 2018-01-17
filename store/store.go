package store

// Task is thing to be done or completed
type Task struct {
	ID     int
	Title  string
	Status string // DOING, PENDING, DONE
}

// Datastore manages a list of tasks stored in memory
type Datastore struct {
	tasks []Task
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
