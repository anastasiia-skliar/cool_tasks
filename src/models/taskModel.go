package models

import (
	"github.com/satori/go.uuid"
	"time"
)

const (
	createTask = "INSERT INTO task (user_id, name, time, created_at, updated_at, desc) VALUES ($1, $1, $3, $4, $5, $6)"
	getTask    = "SELECT * FROM task WHERE id = $1"
	deleteTask = "DELETE FROM task WHERE id = $1"
	getTasks   = "SELECT * FROM task"
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
func CreateTask(task Task) error {
	_, err := db.Exec(createTask, &task.UserID, task.Name, task.Time, task.CreatedAt, task.UpdatedAt, task.Desc)

	return err
}

//GetTask used for getting task from DB
func GetTask(id uuid.UUID) (Task, error) {
	task := Task{}
	err := db.QueryRow(getTask, id).Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)

	return task, err
}

//UpdateTask used fro updating task in DB
func UpdateTask() {

}

//DeleteTask used for deleting task from DB
func DeleteTask(id uuid.UUID) error {
	_, err := db.Exec(deleteTask, id)

	return err
}

//GetTasks used for getting tasks from DB
func GetTasks() ([]Task, error) {
	rows, err := db.Query(getTasks)
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
