package tasksCRUD

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"errors"
	"github.com/satori/go.uuid"
)

type Task struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Time      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Desc      string
}

type successMessage struct {
	Status string `json:"message"`
}

// Start model functions for test

var testID, _ = uuid.FromString("00000000-0000-0000-0000-000000000001")

var tasksArr = []Task{{testID, testID, "Simple Task", time.Now(), time.Now(), time.Now(), "do something"}}
var tasksidArr = Task{testID, testID, "Simple Task", time.Now(), time.Now(), time.Now(), "do something"}

func testGetTaskByID(ID uuid.UUID) (tasksid Task, err error) {

	if ID == tasksidArr.ID {
		tasksid = tasksidArr
	} else {
		err = errors.New("Error")
	}

	return tasksid, err
}
func testGetTasks() (tasks []Task, err error) {
	tasks = tasksArr
	return tasks, err
}
func testDeleteTasks(ID uuid.UUID) (err error) {

	if ID == tasksidArr.ID {
		return nil
	} else {
		err = errors.New("Error")
	}

	return err
}
func testAddTask(task Task) (t Task, err error) {
	return task, err
}

// End model functions for test

func GetTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := testGetTasks()

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

	task, err := testGetTaskByID(taskID)

	if err != nil {
		log.Print(err, " ERROR: Can't get task by ID")
		common.SendError(w, r, 404, "ERROR: Can't get task by ID", err)
		return
	}

	common.RenderJSON(w, r, task)

}

func AddTasks(w http.ResponseWriter, r *http.Request) {

	var newTask Task

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

	task, err := testAddTask(newTask)

	if err != nil {
		log.Print(err, " ERROR: Can't add new task")
		common.SendError(w, r, 400, "ERROR: Can't add new task", err)
		return
	}

	//common.RenderJSON(w, r, successMessage{Status: "Success"})
	common.RenderJSON(w, r, task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := uuid.FromString(params["id"])

	if err != nil {
		log.Print(err, " ERROR: Wrong task ID (can't convert string to uuid)")
		common.SendError(w, r, 400, "ERROR: Wrong task ID (can't convert string to uuid)", err)
		return
	}

	err = testDeleteTasks(taskID)

	if err != nil {
		log.Print(err, " ERROR: Can't delete this task")
		common.SendError(w, r, 404, "ERROR: Can't delete this task", err)
		return
	}

	common.RenderJSON(w, r, successMessage{Status: "Success"})
}
