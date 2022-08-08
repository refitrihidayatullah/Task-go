package entities

// menampung data" task yang diambil dari database
type task struct {
	Id         int64
	TaskDetail string
	Assignee   string
	Status     string
	Deadline   string
}
