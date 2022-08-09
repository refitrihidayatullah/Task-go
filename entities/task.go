package entities

// menampung data" task yang diambil dari database
type Task struct {
	Id         int64
	TaskDetail string `validate:"required" label:"Task Detail"`
	Assignee   string `validate:"required" label:"Assignee"`
	Status     string `validate:"required" label:"Status"`
	Deadline   string `validate:"required" label:"Deadline"`
}
