package model

// Task is thing to be done or completed
type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`   // DOING, PENDING, DONE
	Priority uint8  `json:"priority"` // 1 to 10
}
