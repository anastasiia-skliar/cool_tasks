package models

import (
	"github.com/satori/go.uuid"
	"time"
)

const (
	createTask = "INSERT INTO task (user_id, name, time, created_at, updated_at, desc) VALUES ($1, $1, $3, $4, $5, $6)"
	getTask = "SELECT * FROM task WHERE id = $1"
	deleteTask = "DELETE FROM task WHERE id = $1"
	getTasks = "SELECT * FROM task"
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

func CreateTask(task Task) error {
	_, err := db.Exec(createTask, &task.UserID, task.Name, task.Time, task.CreatedAt, task.UpdatedAt, task.Desc)

	if err != nil {
		return err
	}
	return nil
}

func GetTask(id uuid.UUID) (Task, error) {
	task := Task{}
	err := db.QueryRow(getTask, id).Scan(&task.ID, &task.UserID, &task.Name, &task.Time, &task.CreatedAt, &task.UpdatedAt, &task.Desc)

	if err != nil {
		return task, err
	}

	return task, nil
}

func UpdateTask() {

}

func DeleteTask(id uuid.UUID) error {
	_, err := db.Exec(deleteTask, id)

	if err != nil {
		return err
	}
	return nil
}

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
