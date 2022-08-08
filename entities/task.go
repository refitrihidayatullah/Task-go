package entities

// menampung data" task yang diambil dari database
type Task struct {
	Id         int64
	TaskDetail string
	Assignee   string
	Status     string
	Deadline   string
}
