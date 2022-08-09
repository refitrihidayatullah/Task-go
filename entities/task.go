package entities

// menampung data" task yang diambil dari database
type Task struct {
	Id         int64
	TaskDetail string `validate:"required" label:"Task Detail"`
	Assignee   string `validate:"required"`
	Status     string `validate:"required"`
	Deadline   string `validate:"required"`
}
