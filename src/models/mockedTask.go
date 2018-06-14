package models

import (
	"github.com/satori/go.uuid"
)

func MockedCreateTask(task Task, err error) {
	CreateTask = func(task Task) error {
		return err
	}
}

func MockedGetTask(task Task, err error) {
	GetTask = func(id uuid.UUID) (Task, error) {
		return task, err
	}
}

func MockedDeleteTask(id uuid.UUID, err error)  {
	DeleteTask = func(id uuid.UUID) error {
		return err
	}
}


func MockedGetTasks(task []Task, err error) {
	GetTasks = func() ([]Task, error){
		return task, err
	}
}
