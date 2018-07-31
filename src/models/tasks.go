package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"time"
)

const (
	createTask   = "INSERT INTO tasks (user_id, name, time, created_at, updated_at, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	getTask      = "SELECT * FROM tasks WHERE id = $1"
	deleteTask   = "DELETE FROM tasks WHERE id = $1"
	getTasks     = "SELECT * FROM tasks"
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

//AddTask used for creation task in DB
var AddTask = func(task Task) (Task, error) {
	err := database.DB.QueryRow(createTask, task.UserID, task.Name, task.Time, task.CreatedAt, task.UpdatedAt, task.Desc).Scan(&task.ID)

	return task, err
}

//GetTask used for getting task from DB
var GetTask = func(id uuid.UUID) (Task, error) {
	var task Task
	err := database.DB.QueryRow(getTask, id).Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)

	return task, err
}

//DeleteTask used for deleting task from DB
var DeleteTask = func(id uuid.UUID) error {
	_, err := database.DB.Exec(deleteTask, id)

	return err
}

//GetTasks used for getting tasks from DB
var GetTasks = func() ([]Task, error) {
	rows, err := database.DB.Query(getTasks)
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.Name, &t.Time, &t.CreatedAt, &t.UpdatedAt, &t.Desc); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

//GetUserTasks used for getting all tasks which related to user
var GetUserTasks = func(id uuid.UUID) ([]Task, error) {
	rows, err := database.DB.Query(getUserTasks, id)
	if err != nil {
		return nil, err
	}

	userTasks := make([]Task, 0)

	for rows.Next() {
		task := Task{}
		scanErr := rows.Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)
		if scanErr != nil {
			return nil, scanErr
		}
		userTasks = append(userTasks, task)
	}

	return userTasks, err
}
