//Package tasks implements task handlers
package tasks

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/Nastya-Kruglikova/cool_tasks/src/service/auth"
)

type successCreate struct {
	Status string     `json:"message"`
	Result model.Task `json:"result"`
}

type successChanged struct {
	Status string `json:"message"`
}

type JsonTask struct {
	ID string `json:"id"`
}

//GetTasksHandler gets Tasks from DB
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		return
	}
	if auth.CheckPermission(sessionID, auth.AdminRole, "") == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotAdminResponse, nil)
		return
	}
	tasks, err := model.GetTasks()

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get tasks", err)
		return
	}

	common.RenderJSON(w, r, tasks)
}

//GetTaskHandler gets Task from DB by taskID
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong task ID (can't convert string to uuid)", err)
		return
	}

	task, err := model.GetTask(taskID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get task by ID", err)
		return
	}
	itemOwner, err := model.GetUserByID(task.UserID)
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		return
	}
	if auth.CheckPermission(sessionID, auth.Owner, itemOwner.Login) == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotOwnerResponse, nil)
		return
	}
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get task by ID", err)
		return
	}

	common.RenderJSON(w, r, task)
}

//GetTaskHandler gets Task from DB by taskID
func ChangeStatusHandler(w http.ResponseWriter, r *http.Request) {
	var newTask JsonTask

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTask)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong newTask ID (can't convert string to uuid)", err)
		return
	}

	ID, err := uuid.FromString(newTask.ID)
	if err != nil {
		log.Println(err)
		return
	}

	err = model.ChangeStatus(ID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't change status of this newTask", err)
		return
	}

	task, err := model.GetTask(ID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get task", err)
		return
	}

	itemOwner, err := model.GetUserByID(task.UserID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get user", err)
		return
	}
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		common.SendError(w, r, 400, "ERROR: Can't get cookies", err)
		return
	}

	if auth.CheckPermission(sessionID, auth.Owner, itemOwner.Login) == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotOwnerResponse, nil)
		return
	}
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't change status of this newTask", err)
		return
	}

	common.RenderJSON(w, r, successChanged{Status: "201 Created"})
}

//AddTaskHandler creates and saves Task in DB
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	var newTask model.Task
	var resultTask model.Task

	err := r.ParseForm()

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	timeNow := time.Now()
	userID, err := uuid.FromString(r.Form.Get("user_id"))

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong User ID", err)
		return
	}

	newTask.UserID = userID
	newTask.Name = r.Form.Get("name")
	newTime := r.Form.Get("time")
	newTask.CreatedAt = timeNow
	newTask.UpdatedAt = timeNow
	newTask.Desc = r.Form.Get("desc")

	parsedTime, err := time.Parse(time.UnixDate, newTime)

	if err != nil {
		common.SendUnsupportedMediaType(w, r, "ERROR: Wrong date(can't convert string to int)", err)
		return
	}

	newTask.Time = parsedTime

	resultTask, err = model.AddTask(newTask)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new task", err)
		return
	}

	common.RenderJSON(w, r, successCreate{Status: "201 Created", Result: resultTask})
}

//DeleteTaskHandler deletes Task from DB
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	taskID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong task ID (can't convert string to uuid)", err)
		return
	}

	err = model.DeleteTask(taskID)

	if err != nil {
		common.SendNotFound(w, r, "204 No content", err)
		return
	}

	common.RenderJSON(w, r, nil)
}

//GetUserTasksHandler gets Tasks related to current User
func GetUserTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idUser, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get user", err)
		return
	}

	tasks, err := model.GetUserTasks(idUser)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get user", err)
		return
	}

	itemOwner, err := model.GetUserByID(idUser)
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		return
	}
	if auth.CheckPermission(sessionID, auth.Owner, itemOwner.Login) == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotOwnerResponse, nil)
		return
	}
	common.RenderJSON(w, r, tasks)
}
