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

	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/auth"
)

type successCreate struct {
	Status string     `json:"message"`
	Result model.Task `json:"result"`
}

type successChanged struct {
	Status string `json:"message"`
}

type JsonTaskStatus struct {
	ID string `json:"id"`
}

type JsonTask struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	Desc   string `json:"desc"`
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
	var newTask JsonTaskStatus

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

	var newTask JsonTask
	var resultTask model.Task

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTask)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't decode json from POST Body", err)
		return
	}

	timeNow := time.Now()
	userID, err := uuid.FromString(newTask.UserID)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong User ID", err)
		return
	}

	resultTask.UserID = userID
	resultTask.Name = newTask.Name
	resultTask.CreatedAt = timeNow
	resultTask.UpdatedAt = timeNow
	resultTask.Desc = newTask.Desc

	parsedTime, err := time.Parse(time.UnixDate, newTask.Time)

	if err != nil {
		common.SendUnsupportedMediaType(w, r, "ERROR: Wrong date(can't convert string to time.Time)", err)
		return
	}

	resultTask.Time = parsedTime

	resultTask, err = model.AddTask(resultTask)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new task", err)
		return
	}

	common.RenderJSON(w, r, successCreate{Status: "200 OK", Result: resultTask})
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
