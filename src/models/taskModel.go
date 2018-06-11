package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Task struct {
	ID         uuid.UUID
	User_ID    uuid.UUID
	Name       string
	Time       time.Time
	Created_at time.Time
	Updated_at time.Time
	Desc       string
}

func CreateTask(task Task) error {
	_, err := db.Exec("INSERT INTO task (user_id, name, time, created_at, updated_at, desc) VALUES " +
		"($1, $1, $3, $4, $5, $6)", task.User_ID, task.Name, task.Time, task.Created_at, task.Updated_at, task.Desc)

	if err != nil {
		return err
	}
	return nil
}

func GetTask(id uuid.UUID) (Task, error) {
	task := Task{}
	err := db.QueryRow("SELECT user_id, name, time, created_at, updated_at, desc FROM task WHERE id = $1",
		id).Scan(&task.User_ID, &task.Name, &task.Time, &task.Created_at, &task.Updated_at, &task.Desc)

	if err != nil {
		return task, err
	}

	return task, nil
}

func UpdateTask() {

}

func DeleteTask(id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM task WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT user_id, name, time, created_at, updated_at, desc FROM task")
	if err != nil {

	}

	tasks := make([]Task, 0)

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.User_ID, &t.Name, &t.Time, &t.Created_at, &t.Updated_at, &t.Desc); err != nil {
			return []Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
