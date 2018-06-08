package models

import (
	"time"
)


type Task struct {
	ID         int
	User_ID    int
	Name       string
	Time       time.Time
	Created_at time.Time
	Updated_at time.Time
	Desc       string
}

func createTask()  {

}

func getTask(id int	) (*Task, error) {
	task := &Task{}
	err := db.QueryRow("SELECT user_id, name, time, created_at, updated_at, desc FROM task WHERE id = $1",
		task.ID).Scan(task.User_ID, task.Name, task.Time, task.Created_at, task.Updated_at, task.Desc)

	if err != nil {
		return task, err
	}

	return task, nil
}

func  updateTask()  {

}

func  deleteTask(id int) error {
	_, err := db.Exec("DELETE FROM task WHERE id = $1", id)

	if err != nil{
		return err
	}
	return nil
}




