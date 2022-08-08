package models

import (
	"database/sql"
	"fmt"

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
