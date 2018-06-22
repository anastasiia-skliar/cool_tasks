package models

import (
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"time"
)

const (
	createTask = "INSERT INTO tasks (user_id, name, time, created_at, updated_at, description) VALUES ($1, $2, $3, $4, $5, $6)"
	getTask    = "SELECT * FROM tasks WHERE id = $1"
	deleteTask = "DELETE FROM tasks WHERE id = $1"
	getTasks   = "SELECT * FROM tasks"
	getUserTasks = "SELECT * FROM tasks where user_id = $1"
)

//Task representation in DB
type Task struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Time      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Desc      string
}

//CreateTask used for creation task in DB
var CreateTask = func(task Task) (Task, error) {
	_, err := DB.Exec(createTask, &task.UserID, task.Name, task.Time, task.CreatedAt, task.UpdatedAt, task.Desc)

	return task, err
}

//GetTask used for getting task from DB
var GetTask = func(id uuid.UUID) (Task, error) {
	var task Task
	err := DB.QueryRow(getTask, id).Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)

	return task, err
}

//UpdateTask used fro updating task in DB
var UpdateTask = func() {

}

//DeleteTask used for deleting task from DB
var DeleteTask = func(id uuid.UUID) error {
	_, err := DB.Exec(deleteTask, id)

	return err
}

//GetTasks used for getting tasks from DB
var GetTasks = func() ([]Task, error) {
	rows, err := DB.Query(getTasks)
	if err != nil {
		return []Task{}, err
	}

	tasks := make([]Task, 0)

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.Name, &t.Time, &t.CreatedAt, &t.UpdatedAt, &t.Desc); err != nil {
			return []Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

//UserWithTasks used for getting all tasks which related to user
var GetUserTasks = func(id uuid.UUID) ([]Task, error) {
	rows, err := DB.Query(getUserTasks, id)
	if err != nil {
		return []Task{}, err
	}

	userTasks := make([]Task, 0)

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)

		userTasks = append(userTasks, task)
	}

	return userTasks, err
}
