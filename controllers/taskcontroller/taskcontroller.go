package taskcontroller

import (
	"net/http"
	"text/template"

	"github.com/refitrihidayatullah/task-go/entities"
	"github.com/refitrihidayatullah/task-go/models"
)

// memanggil fuc newTaskModel di models taskmodel.go
var taskModel = models.NewTaskModel()

func Index(response http.ResponseWriter, request *http.Request) {

	// panggil task di taskmodel untuk menampilkan di index
	task, _ := taskModel.FindAll()

	data := map[string]interface{}{
		"task": task,
	}

	temp, err := template.ParseFiles("views/task/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/task/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		// folder struct
		var task entities.Task
		//inputan yg diterima dari form ditangkap dan diteruskan di folder stuct entities task
		task.TaskDetail = request.Form.Get("taskDetail")
		task.Assignee = request.Form.Get("Assignee")
		task.Status = request.Form.Get("Status")
		task.Deadline = request.Form.Get("Deadline")

		// fmt.Println(task)
		// memanggil variabel task model diatas
		taskModel.Create(task)
		// jika berhasil disimpan tampilkan berikut ini
		temp, _ := template.ParseFiles("views/task/add.html")
		data := map[string]interface{}{
			"pesan": "Data task berhasil disimpan",
		}
		temp.Execute(response, data)

	}

}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}
