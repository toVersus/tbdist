package model

// Task is thing to be done or completed
type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`   // DOING, PENDING, DONE
	Priority uint8  `json:"priority"` // 1 to 10
}

// Tasks is just a slice of Task
type Tasks []Task

// Len is the number of elements in the collection
func (t Tasks) Len() int {
	return len(t)
}

// Swap swaps the elements with indexes i and j
func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// ByPriority represents types for sorting by priority number
type ByPriority struct {
	Tasks
}

// Less returns true if the element of index i is less than that of index j
func (b ByPriority) Less(i, j int) bool {
	return b.Tasks[i].Priority < b.Tasks[j].Priority
}
