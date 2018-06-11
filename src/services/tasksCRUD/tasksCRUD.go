package tasksCRUD

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"strconv"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"log"
	"errors"
)

type Task struct {
	ID        int
	UserID    int
	Name      string
	Time      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Desc      string
}

// Start model functions for test

var tasksArr = []Task{{1, 1, "Simple Task", time.Now(), time.Now(), time.Now(), "do something"}}
var tasksidArr = Task{1, 1, "Simple Task", time.Now(), time.Now(), time.Now(), "do something"}

func testGetTaskByID(ID int) (tasksid Task, err error) {

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
func testDeleteTasks(ID int) (err error) {
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
	taskID, atoiErr := strconv.Atoi(params["id"])
	task, err := testGetTaskByID(taskID)

	if atoiErr != nil {
		log.Print(atoiErr, " ERROR: Wrong task ID (can't convert string to int)")
		common.SendError(w, r, 400, "ERROR: Wrong task ID (can't convert string to int)", atoiErr)
		return
	}

	if err != nil {
		log.Print(err, " ERROR: Can't get task by ID")
		common.SendError(w, r, 404, "ERROR: Can't get task by ID", err)
		return
	}

	common.RenderJSON(w, r, task)

}

func AddTasks(w http.ResponseWriter, r *http.Request) {

	var newTask Task

	parseFormErr := r.ParseForm()

	if parseFormErr != nil {
		log.Print(parseFormErr, " ERROR: Can't parse POST Body")
		common.SendError(w, r, 400, "ERROR: Can't parse POST Body", parseFormErr)
		return
	}

	timeNow := time.Now()
	userID, atoiErr := strconv.Atoi(r.Form.Get("user_id"))

	if atoiErr != nil {
		log.Print(atoiErr, " ERROR: Wrong user ID (can't convert string to int)")
		common.SendError(w, r, 400, "ERROR: Wrong user ID (can't convert string to int)", atoiErr)
		return
	}

	newTask.UserID = userID
	newTask.Name = r.Form.Get("name")
	newTime := r.Form.Get("time")
	newTask.CreatedAt = timeNow
	newTask.UpdatedAt = timeNow //When we create new task, updated time = created time
	newTask.Desc = r.Form.Get("desc")

	parsedTime, parsedErr := time.Parse(time.UnixDate, newTime) //parse time from string to time type

	if parsedErr != nil {
		log.Print(parsedErr, " ERROR: Wrong date (can't convert string to int)")
		common.SendError(w, r, 415, "ERROR: Wrong date(can't convert string to int)", parsedErr)
		return
	}

	newTask.Time = parsedTime

	task, taskErr := testAddTask(newTask)

	if taskErr != nil {
		log.Print(taskErr, " ERROR: Can't add new task")
		common.SendError(w, r, 400, "ERROR: Can't add new task", taskErr)
		return
	}

	common.RenderJSON(w, r, task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Print(err, " ERROR: Wrong task ID (can't convert string to int)")
		common.SendError(w, r, 400, "ERROR: Wrong task ID (can't convert string to int)", err)
		return
	}

	deleteErr := testDeleteTasks(taskID)

	if deleteErr != nil {
		log.Print(deleteErr, " ERROR: Can't delete this task")
		common.SendError(w, r, 404, "ERROR: Can't delete this task", deleteErr)
		return
	}
}
