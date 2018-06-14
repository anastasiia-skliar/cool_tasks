package tasksCRUD

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"github.com/satori/go.uuid"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
)


type successMessage struct {
	Status string `json:"message"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := models.GetTasks()

	if err != nil {
		log.Print(err, " ERROR: Can't get tasks")
		common.SendError(w, r, 404, "ERROR: Can't get tasks", err)
		return
	}

	common.RenderJSON(w, r, tasks)

}

func GetTasksByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := uuid.FromString(params["id"])

	if err != nil {
		log.Print(err, " ERROR: Wrong task ID (can't convert string to uuid)")
		common.SendError(w, r, 400, "ERROR: Wrong task ID (can't convert string to uuid)", err)
		return
	}

	task, err := models.GetTask(taskID)

	if err != nil {
		log.Print(err, " ERROR: Can't get task by ID")
		common.SendError(w, r, 404, "ERROR: Can't get task by ID", err)
		return
	}

	common.RenderJSON(w, r, task)

}

func AddTasks(w http.ResponseWriter, r *http.Request) {

	var newTask models.Task

	err := r.ParseForm()

	if err != nil {
		log.Print(err, " ERROR: Can't parse POST Body")
		common.SendError(w, r, 400, "ERROR: Can't parse POST Body", err)
		return
	}

	timeNow := time.Now()
	userID, err := uuid.FromString(r.Form.Get("user_id"))

	if err != nil {
		log.Print(err, " ERROR: Wrong user ID (can't convert string to uuid)")
		common.SendError(w, r, 400, "ERROR: Wrong user ID (can't convert string to uuid)", err)
		return
	}

	newTask.UserID = userID
	newTask.Name = r.Form.Get("name")
	newTime := r.Form.Get("time")
	newTask.CreatedAt = timeNow
	newTask.UpdatedAt = timeNow //When we create new task, updated time = created time
	newTask.Desc = r.Form.Get("desc")

	parsedTime, err := time.Parse(time.UnixDate, newTime) //parse time from string to time type

	if err != nil {
		log.Print(err, " ERROR: Wrong date (can't convert string to int)")
		common.SendError(w, r, 415, "ERROR: Wrong date(can't convert string to int)", err)
		return
	}

	newTask.Time = parsedTime

	err = models.CreateTask(newTask)

	if err != nil {
		log.Print(err, " ERROR: Can't add new task")
		common.SendError(w, r, 400, "ERROR: Can't add new task", err)
		return
	}

	common.RenderJSON(w, r, successMessage{Status: "Success"})
	//common.RenderJSON(w, r, task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := uuid.FromString(params["id"])

	if err != nil {
		log.Print(err, " ERROR: Wrong task ID (can't convert string to uuid)")
		common.SendError(w, r, 400, "ERROR: Wrong task ID (can't convert string to uuid)", err)
		return
	}

	err = models.DeleteTask(taskID)

	if err != nil {
		log.Print(err, " ERROR: Can't delete this task")
		common.SendError(w, r, 404, "ERROR: Can't delete this task", err)
		return
	}

	common.RenderJSON(w, r, successMessage{Status: "Success"})
}
