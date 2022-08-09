package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/refitrihidayatullah/task-go/config"
	"github.com/refitrihidayatullah/task-go/entities"
)

// create task model -> connection database
type TaskModel struct {
	conn *sql.DB
}

// mengembalikan struct model
func NewTaskModel() *TaskModel {
	// Conn ke database
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	//jika tdk ada err maka return conn
	return &TaskModel{
		conn: conn,
	}
}

// func menampilkan data
func (t *TaskModel) FindAll() ([]entities.Task, error) {
	rows, err := t.conn.Query("SELECT * FROM task")
	if err != nil {
		return []entities.Task{}, err
	}
	//jika tdk ada err tutup con database
	defer rows.Close()

	// create variabel  untuk menampung data task
	var dataTask []entities.Task
	for rows.Next() {
		// var menampung 1 task
		var task entities.Task
		rows.Scan(
			&task.Id,
			&task.TaskDetail,
			&task.Assignee,
			&task.Status,
			&task.Deadline,
		)

		// append task ke data task
		if task.Status == "1" {
			task.Status = "Progress"
		} else if task.Status == "2" {
			task.Status = "On-Progress"
		} else {
			task.Status = "Done"
		}
		// yyyy-mm-dd
		deadline, _ := time.Parse("2006-01-02", task.Deadline)
		//dd-mm-yyyy
		task.Deadline = deadline.Format("02-01-2006")
		dataTask = append(dataTask, task)
	}
	return dataTask, nil

}

// struct method untuk proses menyimpan ke databsase
func (t *TaskModel) Create(task entities.Task) bool {
	result, err := t.conn.Exec("insert into task (task_detail,assignee,status,deadline) values(?,?,?,?)", task.TaskDetail, task.Assignee, task.Status, task.Deadline)

	// jika tdk ada err
	if err != nil {
		fmt.Println(err)
		return false
	}
	//jika tdk ada err check lastinsertid
	lastInsertId, _ := result.LastInsertId()
	// jika lastinsert ada maka maka return true || berhasil tersimpan
	return lastInsertId > 0

}

func (t *TaskModel) Find(id int64, task *entities.Task) error {

	return t.conn.QueryRow("select * from task where id =?", id).Scan(
		&task.Id,
		&task.TaskDetail,
		&task.Assignee,
		&task.Status,
		&task.Deadline)
}

// update
func (t *TaskModel) Update(task entities.Task) error {
	_, err := t.conn.Exec("update task set task_detail=?, assignee=?, status=?, deadline=? where id =?",
		task.TaskDetail, task.Assignee, task.Status, task.Deadline, task.Id)

	if err != nil {
		return err
	}
	return nil
}
